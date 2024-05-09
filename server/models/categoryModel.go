package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string `gorm:"unique;not null"`
	CategoryCode uint   `gorm:"unique;not null"`
}
