package models

import (
	"time"
)

type User struct {
	ID        string    `json:"id,omitempty" `
	Email     string    `json:"email,omitempty" validate:"email"`
	Name  string    `json:"name,omitempty"`
	Password  string    ` json:"password,omitempty"`
	CreatedAt time.Time ` json:"createdAt,omitempty"`
	UpdatedAt time.Time ` json:"updatedAt,omitempty"`
}
