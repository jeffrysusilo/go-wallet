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
