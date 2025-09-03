package controllers

import (
	"net/http"
	"altpanel/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"altpanel/utils"
)

func GetDigiScore(c *gin.Context) {
	var req services.DigiScoreRequest
	// Gin automatically validates with "binding:required"
	if err := c.ShouldBindJSON(&req); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range errs {
				var msg string
				switch fieldErr.Field() {
				case "UserRefNumber":
					msg = "User Reference Number is required"
				case "EmploymentType":
					msg = "Employment Type is required"
				default:
					msg = fieldErr.Error()
				}

				c.JSON(http.StatusBadRequest, gin.H{
					"success":    false,
					"error_code": "04",
					"info":       "Validation Error: " + msg,
				})
				return
			}
		}
	}

	// Call service
	result, err := services.GetDigiScore(c, req)

	if err != nil {
		utils.HandleAppError(c, err)
		return
	}
	responseData := gin.H{
		"encrypted_string": result,
	}
	utils.SuccessResponse(c, responseData)
	// c.JSON(http.StatusOK, resp)
	
}
