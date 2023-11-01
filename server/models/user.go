package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time

	UserName  string `json:"UserName"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}
