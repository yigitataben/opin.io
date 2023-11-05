package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	//TODO: The categories in PostView.vue will be transferred to the database.
}

func main() {
	_, err := config.LoadConfigData()
	if err != nil {
		return
	}

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
