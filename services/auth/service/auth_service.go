package service

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"errors"
	"github.com/jackc/pgx/v5"

	"github.com/jeffrysusilo/go-wallet/services/auth/model"
	"github.com/jeffrysusilo/go-wallet/services/auth/repository"
)

func RegisterUser(email, fullName, password string) error {
	existingUser, err := repository.GetUserByEmail(email)
	if err == nil && existingUser != nil {
		return errors.New("email already registered")
	} else if err != pgx.ErrNoRows {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  string(hashedPassword),
		FullName:  fullName,
		CreatedAt: time.Now(),
	}

	return repository.CreateUser(user)
}
