package models

import (
	"time"
)

type User struct {
	ID        string    `json:"id" omitempty`
	Email     string    `json:"email" omitempty`
	Username  string    `json:"username" omitempty`
	Password  string    ` json:"password" omitempty`
	IsAdmin   bool      `json:"isAdmin" omitempty`
	CreatedAt time.Time ` json:"createdAt" omitempty`
	UpdatedAt time.Time ` json:"updatedAt" omitempty`
}
