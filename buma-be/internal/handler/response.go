package handler

import (
	"github.com/labstack/echo/v4"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func RespondWithJSON(c echo.Context, code int, payload interface{}) error {
	return c.JSON(code, payload)
}

func RespondWithError(c echo.Context, code int, message string) error {
	return RespondWithJSON(c, code, APIResponse{
		Success: false,
		Error:   message,
	})
}

func RespondWithSuccess(c echo.Context, code int, message string, data interface{}) error {
	return RespondWithJSON(c, code, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}
