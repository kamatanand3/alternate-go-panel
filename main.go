package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"altpanel/config"
	// "altpanel/repositories"
	"altpanel/routes"
)

func main() {
	// Connect MongoDB
	config.ConnectDB()

	

	// Router
	r := gin.Default()
	// Register routes (with middleware inside)
	routes.RegisterRoutes(r)

	// Get APP_PORT from env
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Server running on port " + port)
	r.Run(":" + port)
}
