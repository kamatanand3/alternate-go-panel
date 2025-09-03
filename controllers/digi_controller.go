package controllers

import (
	"altpanel/services"
	"altpanel/utils"

	"github.com/gin-gonic/gin"
)

func GetDigiScore(c *gin.Context) {
	var req services.DigiScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, 4)
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.ErrorResponse(c, 4, utils.FormatValidationError(err))
		return
	}
	// Call service
	result, err := services.GetDigiScore(c, req)

	if err != nil {
		utils.HandleAppError(c, err)
		return
	}

	utils.SuccessResponse(c, result)

}
