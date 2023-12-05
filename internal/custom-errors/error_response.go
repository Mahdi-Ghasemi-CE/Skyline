package custom_errors

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type apiError struct {
	Field   string
	Message string
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func messageForTag(tag string, field string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("The %s is required.", field)
	case "email":
		return "The input email is invalid. "
	}
	return ""
}

func ValidationErrorMessage(err error) gin.H {
	var violations validator.ValidationErrors
	if errors.As(err, &violations) {
		out := make([]apiError, len(violations))
		for i, e := range violations {
			out[i] = apiError{e.Field(), messageForTag(e.Tag(), e.Field())}
		}

		return gin.H{"error": out}
	}
	return nil
}
