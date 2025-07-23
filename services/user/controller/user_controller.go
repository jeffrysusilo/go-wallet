package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jeffrysusilo/go-wallet/services/user/config"
	"github.com/jeffrysusilo/go-wallet/services/user/model"
	"github.com/google/uuid"
)

func GetUserProfileByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var user model.User
	query := `SELECT id, full_name, email, role FROM users WHERE id = $1`
	err = config.DB.QueryRow(context.Background(), query, userID).Scan(
		&user.ID, &user.FullName, &user.Email, &user.Role,
	)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func GetMyProfile(c *fiber.Ctx) error {
	userIDStr := c.Locals("user_id")
	if userIDStr == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var user model.User
	query := `SELECT id, full_name, email, role FROM users WHERE id = $1`
	err = config.DB.QueryRow(context.Background(), query, userID).Scan(
		&user.ID, &user.FullName, &user.Email, &user.Role,
	)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func UpdateMyProfile(c *fiber.Ctx) error {
	userIDStr := c.Locals("user_id")
	if userIDStr == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var body struct {
		FullName string `json:"full_name"`
		Email    string `json:"email"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if body.FullName == "" || body.Email == "" {
		return c.Status(400).JSON(fiber.Map{"error": "full_name and email are required"})
	}

	query := `UPDATE users SET full_name = $1, email = $2 WHERE id = $3`
	_, err = config.DB.Exec(context.Background(), query, body.FullName, body.Email, userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	return c.JSON(fiber.Map{"message": "Profile updated successfully"})
}

func GetAllUsers(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT id, full_name, email, role FROM users")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch users"})
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.FullName, &user.Email, &user.Role)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error scanning user"})
		}
		users = append(users, user)
	}

	return c.JSON(users)
}

