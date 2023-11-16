package routes

import (
	"github.com/gofiber/fiber/v2"
	"server/database"
)

func GetAllCategories(c *fiber.Ctx) error {
	categories, err := database.GetCategories()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error retrieving categories"})
	}

	return c.JSON(categories)
}
