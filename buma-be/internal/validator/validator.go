package validator

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errMsgs []string
			for _, vErr := range validationErrors {
				fieldName := toSnakeCase(vErr.Field())
				switch vErr.Tag() {
				case "required":
					errMsgs = append(errMsgs, fmt.Sprintf("%s is required", fieldName))
				case "email":
					errMsgs = append(errMsgs, fmt.Sprintf("%s must be a valid email address", fieldName))
				case "min":
					errMsgs = append(errMsgs, fmt.Sprintf("%s must be at least %s characters long", fieldName, vErr.Param()))
				default:
					errMsgs = append(errMsgs, fmt.Sprintf("%s failed validation for tag '%s'", fieldName, vErr.Tag()))
				}
			}
			return echo.NewHTTPError(http.StatusBadRequest, strings.Join(errMsgs, ", "))
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func toSnakeCase(str string) string {
	var res strings.Builder
	for i, r := range str {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				res.WriteRune('_')
			}
			res.WriteRune(r + 32)
		} else {
			res.WriteRune(r)
		}
	}
	return res.String()
}
