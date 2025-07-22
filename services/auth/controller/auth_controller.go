package controller

import (
	"time"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/jeffrysusilo/go-wallet/services/auth/repository"

	"github.com/jeffrysusilo/go-wallet/services/auth/service"

	"github.com/google/uuid"
	"context"
	"github.com/jeffrysusilo/go-wallet/services/auth/config"
	"github.com/jeffrysusilo/go-wallet/services/auth/model"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

func Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if req.Email == "" || req.Password == "" || req.FullName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Missing required fields"})
	}

	err := service.RegisterUser(req.Email, req.FullName, req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	secret := os.Getenv("JWT_SECRET")
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	refreshToken := uuid.NewString()
	updateQuery := `UPDATE users SET refresh_token = $1 WHERE id = $2`
	_, err = config.DB.Exec(context.Background(), updateQuery, refreshToken, user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to store refresh token"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		MaxAge:   86400 * 7, // 7 hari
		Secure:   true,     // true jika pakai HTTPS di prod
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"token": signedToken,
	})
}

func RefreshToken(c *fiber.Ctx) error {
	cookie := c.Cookies("refresh_token")
	if cookie == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Missing refresh token"})
	}

	var user model.User
	query := `SELECT id, email, role FROM users WHERE refresh_token = $1`
	err := config.DB.QueryRow(context.Background(), query, cookie).Scan(&user.ID, &user.Email, &user.Role)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid refresh token"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role, 
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": signedToken})
}

func Logout(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken != "" {
		// Hapus dari DB
		_, _ = config.DB.Exec(context.Background(), `UPDATE users SET refresh_token = NULL WHERE refresh_token = $1`, refreshToken)
	}

	// Hapus cookie
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"message": "Logged out"})
}
