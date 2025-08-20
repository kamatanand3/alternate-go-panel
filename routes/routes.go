package routes

import (
	"github.com/gin-gonic/gin"
	"altpanel/controllers"
	"altpanel/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	r.Use(gin.Recovery())    // handles panic
	r.Use(middleware.ApiHandler()) // custom header validation middleware
	r.Use(middleware.Logger()) // log request + response


	v1 := r.Group("/api/v1")
	{
		// v1.POST("/users", controllers.CreateUser)
		// v1.GET("/users", controllers.GetAllUsers)
		// v1.GET("/users/:id", controllers.GetUserByID)
		// v1.PUT("/users/:id", controllers.UpdateUser)
		// v1.DELETE("/users/:id", controllers.DeleteUser)
		v1.GET("/configs", controllers.GetAllConfig)
	}	
}