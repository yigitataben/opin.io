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
)

func setupRoutes(app *fiber.App) {
	// User endpoints:
	app.Post("/user", routes.CreateUser)
	app.Get("/user", routes.GetAllUsers)
	app.Get("/user/:id", routes.GetUserByID)
	app.Put("/user/:id", routes.UpdateUser)
	app.Delete("/user/:id", routes.DeleteUser)

	// Post endpoints:
	app.Post("/post", routes.CreatePost)
	app.Get("/post", routes.GetAllPosts)
	app.Get("/post/:id", routes.GetPostByID)
	app.Put("/post/:id", routes.UpdatePost)
	app.Delete("/post/:id", routes.DeletePost)

	// Post categories:
	app.Get("/categories", routes.GetCategories)
	app.Post("/categories", routes.CreateCategory)
	app.Post("/log-in", routes.LoginUser)
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
		AllowHeaders:     "Content-Type",
	}))

	setupRoutes(app)

	port := viper.GetInt("server.port")
	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Failed to start the server: %s", err)
	}
}
