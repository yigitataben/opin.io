package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"log"
	"server/config"
	"server/database"
	"server/routes"
	"server/session"
)

func setupRoutes(app *fiber.App) {
	app.Use(authMiddleware)

	// User endpoints:
	app.Post("/user", routes.CreateUser)
	app.Get("/user", routes.GetAllUsers)
	app.Get("/user/:id", routes.GetUserByID)
	app.Put("/user/:id", routes.UpdateUser)
	app.Delete("/user/:id", routes.DeleteUser)
	app.Post("/log-in", routes.LoginUser)
	app.Post("/log-out", routes.LogoutUser)

	// Post endpoints:
	app.Post("/post", routes.CreatePost)
	app.Get("/post", routes.GetAllPosts)
	app.Get("/post/:id", routes.GetPostByID)
	app.Put("/post/:id", routes.UpdatePost)
	app.Delete("/post/:id", routes.DeletePost)

	// Post categories:
	app.Get("/categories", routes.GetCategories)
	app.Post("/categories", routes.CreateCategory)
}

func authMiddleware(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing JWT token",
		})
	}

	_, err := session.VerifyToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid JWT token",
		})
	}

	return c.Next()
}

func main() {
	if _, err := config.LoadConfigData(); err != nil {
		log.Fatalf("Failed to load configuration data: %s", err)
	}

	if err := database.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect to the database: %s", err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE",
	}))

	setupRoutes(app)

	port := viper.GetInt("server.port")
	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Failed to start the server: %s", err)
	}
}
