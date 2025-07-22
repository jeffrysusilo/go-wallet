package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Jangan expose password
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}
