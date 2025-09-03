package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ResponseCode struct {
	RequestID    string      `json:"request_id"`
	Success      bool        `json:"success"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
	HTTPCode     int         `json:"-"`
}

// GetResponseCode returns standardized response codes similar to Laravel ResponseCodeTrait
func GetResponseCode(c *gin.Context, code int) ResponseCode {
	responseCodes := map[int]ResponseCode{
		1: {
			RequestID:    GetRequestIDFromContext(c),
			Success:      true,
			ResponseCode: 0,
			Message:      "Success",
			HTTPCode:     200,
		},
		2: {
			RequestID:    GetRequestIDFromContext(c),
			Success:      false,
			ResponseCode: 0,
			Message:      "Success",
			HTTPCode:     200,
		},
		101: {
			RequestID:    GetRequestIDFromContext(c),
			Success:      false,
			ResponseCode: 101,
			Message:      "Validation failed",
			HTTPCode:     400,
		},
		201: {
			RequestID:    GetRequestIDFromContext(c),
			Success:      false,
			ResponseCode: 201,
			Message:      "Invalid Key Id OR Key Secret",
			HTTPCode:     401,
		},
		203: {
			RequestID:    GetRequestIDFromContext(c),
			Success:      false,
			ResponseCode: 203,
			Message:      "Validation failed",
			HTTPCode:     401,
		},
		3: {
			RequestID:    GetRequestIDFromContext(c),
			Success:      false,
			ResponseCode: 2,
			Message:      "Not found",
			HTTPCode:     404,
		},
		102: {
			RequestID:    GetRequestIDFromContext(c),
			Success:      false,
			ResponseCode: 102,
			Message:      "Internal server error",
			HTTPCode:     500,
		},
		429: {
			RequestID:    GetRequestIDFromContext(c),
			Success:      false,
			ResponseCode: 429,
			Message:      "Too many request, please try after sometime",
			HTTPCode:     429,
		},
	}

	if response, exists := responseCodes[code]; exists {
		return response
	}

	// Default response
	return ResponseCode{
		RequestID:    GetRequestIDFromContext(c),
		Success:      false,
		ResponseCode: 999,
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
func ErrorResponse(c *gin.Context, code int, message string) {
	response := GetResponseCode(c, code)
	if message != "" {
		response.Message = message
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
