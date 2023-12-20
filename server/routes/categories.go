package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"server/database"
	"server/models"
	"time"
)

type CategoryResponse struct {
	ID           uint      `json:"id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
}

func CreateResponseCategoryModel(categoryModel models.Category) CategoryResponse {
	return CategoryResponse{
		ID:           categoryModel.CategoryID,
		CategoryName: categoryModel.CategoryName,
		CreatedAt:    categoryModel.CreatedAt,
	}
}

func CreateCategory(c *fiber.Ctx) error {
	var categories []models.Category

	if err := c.BodyParser(&categories); err != nil {
		fmt.Println("Error parsing request body:", err)
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	for i := range categories {
		categories[i].CategoryID = 0

		result := database.Database.DB.Create(&categories[i])
		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
		}
	}

	responseCategories := make([]CategoryResponse, len(categories))
	for i, category := range categories {
		responseCategories[i] = CreateResponseCategoryModel(category)
	}

	return c.Status(201).JSON(responseCategories)
}

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category

	database.Database.DB.Find(&categories)
	responseCategories := make([]CategoryResponse, len(categories))

	for i, category := range categories {
		responseCategories[i] = CreateResponseCategoryModel(category)
	}

	return c.Status(200).JSON(responseCategories)
}

func GetCategoryByID(c *fiber.Ctx) error {
	categoryID := c.Params("id")
	var category models.Category

	result := database.Database.DB.First(&category, categoryID)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Category not found"})
	}

	responseCategory := CreateResponseCategoryModel(category)

	return c.Status(200).JSON(responseCategory)
}

func UpdateCategory(c *fiber.Ctx) error {
	categoryID := c.Params("id")
	category := new(models.Category)

	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result := database.Database.DB.Model(&models.Category{}).Where("category_id = ?", categoryID).Updates(category)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	responseCategory := CreateResponseCategoryModel(*category)

	return c.Status(200).JSON(responseCategory)
}

func DeleteCategory(c *fiber.Ctx) error {
	categoryID := c.Params("id")

	result := database.Database.DB.Delete(&models.Category{}, categoryID)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Category deleted successfully"})
}
