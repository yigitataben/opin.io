package models

import "time"

// User struct represents a user entity.
type User struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`

	UserName     string `json:"user_name"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	EmailAddress string `json:"email_address"`
	UserPassword string `json:"user_password"`
}
