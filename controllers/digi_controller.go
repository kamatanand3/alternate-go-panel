package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	service "altpanel/services"
)


func GetDigiScore(c *gin.Context) {
	users, err := service.GetDigiScore()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}



// var digiService *services.DigiService

// // InitDigiController initializes DigiService with repositories
// func InitDigiController(service *services.DigiService) {
// 	// digiService = service
// }


// Handler function
// func GetDigiScore(c *gin.Context) {
// 	var req services.DigiScoreRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success":    false,
// 			"error_code": "04",
// 			"info":       "Invalid request data",
// 		})
// 		return
// 	}

// 	// Call service
// 	resp, code := digiService.GetDigiScore(req)
// 	c.JSON(code, resp)
// }
