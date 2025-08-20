package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "altpanel/services"
)


func GetAllConfig(c *gin.Context) {
	users, err := service.GetAllConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

