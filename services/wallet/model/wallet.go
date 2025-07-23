package model

import "github.com/google/uuid"

type Wallet struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Balance int64     `json:"balance"`
}
