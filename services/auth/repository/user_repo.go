package repository

import (
	"context"
	// "time"

	"github.com/jeffrysusilo/go-wallet/services/auth/config"
	"github.com/jeffrysusilo/go-wallet/services/auth/model"
)

func CreateUser(user *model.User) error {
	query := `
		INSERT INTO users (id, email, password_hash, full_name, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := config.DB.Exec(
		context.Background(),
		query,
		user.ID, user.Email, user.Password, user.FullName, user.CreatedAt,
	)
	return err
}

func GetUserByEmail(email string) (*model.User, error) {
	query := `SELECT id, email, password_hash, full_name, created_at FROM users WHERE email = $1`
	row := config.DB.QueryRow(context.Background(), query, email)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.FullName, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

