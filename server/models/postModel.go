package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content      string `gorm:"unique;not null"`
	CategoryName string `gorm:"unique;not null;foreignkey:CategoryName"`
	CategoryCode uint   `gorm:"unique;not null;foreignkey:CategoryCode"`
	UserID       uint   `gorm:"unique;not null;foreignkey:UserID"`
}
