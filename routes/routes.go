package routes

import (
	"altpanel/controllers"
	"altpanel/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.Use(gin.Recovery())          // handles panic
	r.Use(middleware.ApiHandler()) // custom header validation middleware
	// r.Use(middleware.Logger()) // log request + response

	v1 := r.Group("/api/v1")
	{
		v1.POST("/digiscore", controllers.GetDigiScore)
	}
}
