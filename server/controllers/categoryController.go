package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yigitataben/go-jwt/initializers"
	"github.com/yigitataben/go-jwt/models"
	"net/http"
)

func CreateCategories(c *gin.Context) {
	type CategoryInput struct {
		CategoryName string `json:"category_name"`
	}
	var categoriesInput []CategoryInput
	if err := c.BindJSON(&categoriesInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}
	var categories []models.Category
	for i, input := range categoriesInput {
		category := models.Category{
			CategoryName: input.CategoryName,
			CategoryCode: uint(i + 1),
		}
		categories = append(categories, category)
	}
	result := initializers.DB.Create(&categories)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	result := initializers.DB.Order("id asc").Find(&categories)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve categories",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

/*
func GetCategory(c *gin.Context) {
	ID := c.Param("id")
	var categories models.Category
	result := initializers.DB.First(&categories, ID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Category not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func DeleteCategory(c *gin.Context) {
	ID := c.Param("id")
	var categories models.Category
	result := initializers.DB.First(&categories, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	result = initializers.DB.Unscoped().DeleteUser(&categories)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete category",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
	})
}
*/
