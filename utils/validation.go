package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func FormatValidationError(err error) string {
	var sb strings.Builder
	for _, e := range err.(validator.ValidationErrors) {
		// fmt.Printf("Field: %s, Tag: %s, ActualTag: %s, Kind: %s, Type: %s, Value: %v\n",
		sb.WriteString(fmt.Sprintf("Validation Error: '%s' is missing which is '%s'", e.Field(), e.Tag()))
	}
	return sb.String()
}
