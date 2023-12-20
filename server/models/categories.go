package models

import "time"

// Category struct represents a category entity.
type Category struct {
	CategoryID uint      `json:"category_id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"created_at"`

	CategoryName string `json:"category_name"`
}
