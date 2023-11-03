package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"server/database"
	"server/routes"
)

func setupRoutes(app *fiber.App) {
	// User endpoints:
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetAllUsers)
	app.Get("/api/users/:id", routes.GetUserByID)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// Post endpoints:
	app.Post("/zurna", routes.CreatePost)
	app.Get("/api/posts", routes.GetAllPosts)
	app.Get("/api/posts/:id", routes.GetPostByID)
	app.Put("/api/posts/:id", routes.UpdatePost)
	app.Delete("/api/posts/:id", routes.DeletePost)
}

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE",
	}))
	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
