package initializers

import "github.com/yigitataben/go-jwt/models"

func SyncDB() {
	err := DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Category{})
	if err != nil {
		return
	}
}
