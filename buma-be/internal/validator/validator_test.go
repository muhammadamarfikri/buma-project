package validator_test

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"

	"buma-be/internal/domain"
	"buma-be/internal/validator"
)

func TestCustomValidator(t *testing.T) {
	v := validator.NewCustomValidator()

	t.Run("Valid Officer Request", func(t *testing.T) {
		req := domain.CreateOfficerRequest{
			FullName:    "Budi Santoso",
			Email:       "budi@buma.co.id",
			PairingTeam: "Desk 01",
		}
		if err := v.Validate(&req); err != nil {
			t.Errorf("Expected no validation error for valid request, got: %v", err)
		}
	})

	t.Run("Missing Required Fields", func(t *testing.T) {
		req := domain.CreateOfficerRequest{
			Email: "invalid-email",
		}
		err := v.Validate(&req)
		if err == nil {
			t.Fatalf("Expected validation error for missing required fields, got nil")
		}

		he, ok := err.(*echo.HTTPError)
		if !ok || he.Code != http.StatusBadRequest {
			t.Errorf("Expected HTTP 400 Bad Request error, got: %v", err)
		}
	})

	t.Run("Valid Commitment Request", func(t *testing.T) {
		req := domain.CreateCommitmentRequest{
			AccountNo: "1029384756",
			Reason:    "Setor cicilan",
		}
		if err := v.Validate(&req); err != nil {
			t.Errorf("Expected no validation error for valid commitment request, got: %v", err)
		}
	})
}
