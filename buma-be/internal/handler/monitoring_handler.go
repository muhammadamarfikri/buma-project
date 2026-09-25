package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	"buma-be/internal/domain"
	"buma-be/internal/usecase"
)

type MonitoringHandler struct {
	usecase usecase.MonitoringUsecase
}

func NewMonitoringHandler(u usecase.MonitoringUsecase) *MonitoringHandler {
	return &MonitoringHandler{usecase: u}
}

// GetMonitoring handles GET /api/v1/monitoring
func (h *MonitoringHandler) GetMonitoring(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	filter := domain.DebtorFilter{
		Page:        page,
		Limit:       limit,
		Search:      strings.TrimSpace(c.QueryParam("search")),
		Tier:        strings.TrimSpace(c.QueryParam("tier")),
		OfficerID:   strings.TrimSpace(c.QueryParam("manager_id")), // supports manager_id or officer_id
		OfficerName: strings.TrimSpace(c.QueryParam("pengelola")),
		Product:     strings.TrimSpace(c.QueryParam("produk")),
		SortBy:      strings.TrimSpace(c.QueryParam("sort_by")),
		SortDir:     strings.TrimSpace(c.QueryParam("sort_dir")),
	}

	if filter.OfficerID == "" {
		filter.OfficerID = strings.TrimSpace(c.QueryParam("officer_id"))
	}

	resp, err := h.usecase.GetMonitoring(c.Request().Context(), filter)
	if err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	return RespondWithJSON(c, http.StatusOK, resp)
}

// GetMonitoringByAccountNo handles GET /api/v1/monitoring/:account_no
func (h *MonitoringHandler) GetMonitoringByAccountNo(c echo.Context) error {
	accountNo := c.Param("account_no")
	if accountNo == "" {
		return RespondWithError(c, http.StatusBadRequest, "account_no parameter is required")
	}

	debtor, err := h.usecase.GetDebtorByAccountNo(c.Request().Context(), accountNo)
	if err != nil {
		return RespondWithError(c, http.StatusNotFound, err.Error())
	}

	return RespondWithSuccess(c, http.StatusOK, "Debtor detail retrieved successfully", debtor)
}
