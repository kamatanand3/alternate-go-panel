package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ResponseCode struct {
	RequestID    string      `json:"request_id"`
	Success      bool        `json:"success"`
	ResponseCode string      `json:"error_code"`
	Message      string      `json:"info"`
	Data         interface{} `json:"data,omitempty"`
	HTTPCode     int         `json:"-"`
}

// GetResponseCode returns standardized response codes similar to Laravel ResponseCodeTrait
func GetResponseCode(c *gin.Context, code int) ResponseCode {
	responseCodes := map[int]ResponseCode{
		1: {
			RequestID:    GetRequestIDFromContext(c),
			Success:      true,
			ResponseCode: "0",
			Message:      "Success",
			HTTPCode:     200,
		},
		2: {
			RequestID:    GetRequestIDFromContext(c),
			Success:      false,
			ResponseCode: "0",
			Message:      "Success",
			HTTPCode:     200,
		},
		4: {
			RequestID:    GetRequestIDFromContext(c),
			Success:      false,
			ResponseCode: "04",
			Message:      "Invalid json",
			HTTPCode:     400,
		},
	}

	if response, exists := responseCodes[code]; exists {
		return response
	}

	// Default response
	return ResponseCode{
		RequestID:    GetRequestIDFromContext(c),
		Success:      false,
		ResponseCode: "999",
		Message:      "Unknown error",
		HTTPCode:     500,
	}
}

// Response sends a standardized JSON response
func Response(c *gin.Context, response ResponseCode) {
	httpCode := response.HTTPCode
	if httpCode == 0 {
		httpCode = 200
	}
	resp := ResponseCode{
		RequestID:    response.RequestID,
		Success:      response.Success,
		ResponseCode: response.ResponseCode,
		Message:      response.Message,
		Data:         response.Data,
	}
	c.JSON(httpCode, resp)

}

// SuccessResponse sends a success response with data
func SuccessResponse(c *gin.Context, data interface{}) {
	response := GetResponseCode(c, 1)
	response.Data = data
	Response(c, response)
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, code int, message ...string) {
	response := GetResponseCode(c, code)
	if len(message) > 0 && message[0] != "" {
		response.Message = message[0]
	}
	Response(c, response)
}

// ValidationErrorResponse sends a validation error response
func ValidationErrorResponse(c *gin.Context, message string) {
	response := GetResponseCode(c, 101)
	if message != "" {
		response.Message = message
		c.Error(fmt.Errorf(message))
	}
	Response(c, response)
}

func ServiceFailedErrorResponse(c *gin.Context, message string) {
	response := GetResponseCode(c, 203)
	if message != "" {
		response.Message = message
		c.Error(fmt.Errorf(message))
	}
	Response(c, response)
}
func InvalidKeyErrorResponse(c *gin.Context, message string) {
	response := GetResponseCode(c, 201)
	if message != "" {
		response.Message = message
		c.Error(fmt.Errorf(message))
	}
	Response(c, response)
}

// NotFoundResponse sends a not found response
func NotFoundResponse(c *gin.Context, message string) {
	response := GetResponseCode(c, 3)
	if message != "" {
		response.Message = message
	}
	Response(c, response)
}

// InternalErrorResponse sends an internal server error response
func InternalErrorResponse(c *gin.Context, message string) {
	response := GetResponseCode(c, 102)
	if message != "" {
		response.Message = message
		c.Error(fmt.Errorf(message))
	}
	Response(c, response)
}

// InternalErrorResponse sends an internal server error response
func ErrorTooManyRequestsResponse(c *gin.Context, message string) {
	response := GetResponseCode(c, 429)
	if message != "" {
		response.Message = message
		c.Error(fmt.Errorf(message))
	}
	Response(c, response)
}
