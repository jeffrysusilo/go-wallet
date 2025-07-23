package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jeffrysusilo/go-wallet/services/wallet/config"
	"github.com/jeffrysusilo/go-wallet/services/wallet/model"
)

func GetBalance(c *fiber.Ctx) error {
	userIDStr := c.Locals("user_id")
	if userIDStr == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var wallet model.Wallet
	query := `SELECT id, user_id, balance FROM wallets WHERE user_id = $1`
	err = config.DB.QueryRow(context.Background(), query, userID).Scan(
		&wallet.ID, &wallet.UserID, &wallet.Balance,
	)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Wallet not found"})
	}

	return c.JSON(fiber.Map{
		"wallet_id": wallet.ID,
		"balance":   wallet.Balance,
	})
}
