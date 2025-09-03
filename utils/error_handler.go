package utils

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func HandleAppError(c *gin.Context, err error) {
	var appErr AppError
	if errors.As(err, &appErr) {
		var msg string
		switch {
		case appErr.Message != "" && appErr.Err != nil:
			msg = appErr.Message + " (" + appErr.Err.Error() + ")"
		case appErr.Message != "":
			msg = appErr.Message
		case appErr.Err != nil:
			msg = appErr.Err.Error()
		default:
			msg = string(appErr.Type)
		}

		switch appErr.Type {
		case ErrorTypeValidation:
			ValidationErrorResponse(c, msg)
		case ErrorTypeInternal:
			ServiceFailedErrorResponse(c, msg)
		case ErrorTypeEncryption:
			InternalErrorResponse(c, msg)
		case ErrorTypeInvalidInput:
			InvalidKeyErrorResponse(c, msg)
		case ErrorTooManyRequests:
			ErrorTooManyRequestsResponse(c, msg)
		default:
			InternalErrorResponse(c, "An internal error occurred")
			log.Printf("Internal error: %v", err)
		}
		return
	}
	InternalErrorResponse(c, "An unexpected error occurred")
	log.Printf("Unexpected error: %v", err)
}
