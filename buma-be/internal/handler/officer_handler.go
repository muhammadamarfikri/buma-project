package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"buma-be/internal/domain"
	"buma-be/internal/usecase"
)

type OfficerHandler struct {
	usecase usecase.OfficerUsecase
}

func NewOfficerHandler(u usecase.OfficerUsecase) *OfficerHandler {
	return &OfficerHandler{usecase: u}
}

// CreateOfficer handles POST /api/v1/officers
func (h *OfficerHandler) CreateOfficer(c echo.Context) error {
	var req domain.CreateOfficerRequest
	if err := c.Bind(&req); err != nil {
		return RespondWithError(c, http.StatusBadRequest, "Invalid JSON request body")
	}

	if err := c.Validate(&req); err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	officer, err := h.usecase.CreateOfficer(c.Request().Context(), req)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	return RespondWithSuccess(c, http.StatusCreated, "Officer added successfully", officer)
}

// GetAllOfficers handles GET /api/v1/officers
func (h *OfficerHandler) GetAllOfficers(c echo.Context) error {
	officers, err := h.usecase.GetAllOfficers(c.Request().Context())
	if err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	return RespondWithSuccess(c, http.StatusOK, "Officers retrieved successfully", officers)
}

// GetOfficerByID handles GET /api/v1/officers/:officer_id
func (h *OfficerHandler) GetOfficerByID(c echo.Context) error {
	officerID := c.Param("officer_id")
	if officerID == "" {
		return RespondWithError(c, http.StatusBadRequest, "officer_id parameter is required")
	}

	officer, err := h.usecase.GetOfficerByID(c.Request().Context(), officerID)
	if err != nil {
		return RespondWithError(c, http.StatusNotFound, err.Error())
	}

	return RespondWithSuccess(c, http.StatusOK, "Officer details retrieved successfully", officer)
}
