package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"altpanel/config"
	"altpanel/controllers"
	"altpanel/repositories"
)

func main() {
	// Connect MongoDB
	config.ConnectDB()

	// Init repo collections
	// repositories.InitUserRepository()
	repositories.InitConfigRepository()

	// Router
	r := gin.Default()

	// Routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetUserByID)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.GET("/configs", controllers.GetAllConfig)

	// Get APP_PORT from env
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Server running on port " + port)
	r.Run(":" + port)
}
