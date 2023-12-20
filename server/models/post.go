package models

import "time"

// Post struct represents a post entity.
type Post struct {
	PostID    uint      `json:"post_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`

	Content string `json:"content"`

	UserID   uint   `json:"user_id" gorm:"index"`
	UserName string `json:"user_name" gorm:"index"`

	CategoryName string `json:"category_name" gorm:"index"`
	CategoryID   uint   `json:"category_id" gorm:"index"`
}
