package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yigitataben/go-jwt/controllers"
	"github.com/yigitataben/go-jwt/initializers"
	"github.com/yigitataben/go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	// CORS middleware
	r.Use(CORSMiddleware())

	// Authorization endpoints:
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	// User endpoints:
	r.GET("/user", controllers.GetAllUsers)
	r.GET("/user/:id", controllers.GetUserByID)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	// Category endpoints:
	r.POST("/categories", controllers.CreateCategories)
	r.GET("/categories", controllers.GetAllCategories)

	// Post endpoints:
	r.POST("/post", controllers.CreatePost)
	r.GET("/post", controllers.GetAllPosts)
	r.GET("/post/:id", controllers.GetPostByID)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)

	r.Run()
}

// CORSMiddleware adds the necessary CORS headers to the response.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
