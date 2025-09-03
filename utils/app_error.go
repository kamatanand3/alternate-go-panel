package utils

import (
	"fmt"
)

type ErrorType string

const (
    ErrorTypeInvalidInput   ErrorType = "INVALID_INPUT"
    ErrorTypeAuthentication ErrorType = "AUTHENTICATION"
    ErrorTypeEncryption     ErrorType = "ENCRYPTION"
    ErrorTypeInternal       ErrorType = "INTERNAL"
    ErrorTypeValidation     ErrorType = "VALIDATION"
    ErrorTooManyRequests    ErrorType = "TOO_MANY_REQUESTS"
)

type AppError struct {
    Type    ErrorType
    Message string
    Err     error
}

func (e AppError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("%s: %s (%s)", e.Type, e.Message, e.Err.Error())
    }
    return fmt.Sprintf("%s: %s", e.Type, e.Message)
}