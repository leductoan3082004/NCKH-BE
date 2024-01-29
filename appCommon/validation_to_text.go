package appCommon

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ValidationErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", ToSnake(e.Field()))
	case "max":
		return fmt.Sprintf("%s cannot be longer than %s", ToSnake(e.Field()), e.Param())
	case "min":
		return fmt.Sprintf("%s must be longer than %s", ToSnake(e.Field()), e.Param())
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("%s must be %s characters long", ToSnake(e.Field()), e.Param())
	}
	return fmt.Sprintf("%s is not valid", ToSnake(e.Field()))
}
