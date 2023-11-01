package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"server/database"
	"server/routes"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString(("API is working!"))
}
func setupRoutes(app *fiber.App) {
	// User endpoints:
	app.Get("/api", healthCheck)
	app.Post("api/users", routes.CreateUser)
	app.Get("api/users", routes.GetAllUsers)
	app.Get("/api/users/:id", routes.GetUserByID)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// Post endpoints:
	app.Post("api/posts", routes.CreatePost)
	app.Get("api/posts", routes.GetAllPosts)
	app.Get("/api/posts/:id", routes.GetPostByID)
	app.Put("/api/posts/:id", routes.UpdatePost)
	app.Delete("/api/posts/:id", routes.DeletePost)
}
func main() {
	database.ConnectDB()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
