package models

import "time"

type Post struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time

	PostCategory string `json:"PostCategory"`
	PostTitle    string `json:"PostTitle"`
	PostBody     string `json:"PostBody"`
}
