package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	"buma-be/internal/domain"
	"buma-be/internal/usecase"
)

type CommitmentHandler struct {
	usecase usecase.CommitmentUsecase
}

func NewCommitmentHandler(u usecase.CommitmentUsecase) *CommitmentHandler {
	return &CommitmentHandler{usecase: u}
}

// CreateCommitment handles POST /api/v1/commitments
func (h *CommitmentHandler) CreateCommitment(c echo.Context) error {
	var req domain.CreateCommitmentRequest
	if err := c.Bind(&req); err != nil {
		return RespondWithError(c, http.StatusBadRequest, "Invalid JSON request body")
	}

	if err := c.Validate(&req); err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	commitment, err := h.usecase.CreateCommitment(c.Request().Context(), req)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	return RespondWithSuccess(c, http.StatusCreated, "Commitment recorded successfully", commitment)
}

// GetAllCommitments handles GET /api/v1/commitments
func (h *CommitmentHandler) GetAllCommitments(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	filter := domain.CommitmentFilter{
		Page:      page,
		Limit:     limit,
		AccountNo: strings.TrimSpace(c.QueryParam("account_no")),
		Status:    strings.TrimSpace(c.QueryParam("status")),
		Search:    strings.TrimSpace(c.QueryParam("search")),
		Tier:      strings.TrimSpace(c.QueryParam("tier")),
		Pengelola: strings.TrimSpace(c.QueryParam("pengelola")),
		Product:   strings.TrimSpace(c.QueryParam("product")),
		SortBy:    strings.TrimSpace(c.QueryParam("sort_by")),
		SortDir:   strings.TrimSpace(c.QueryParam("sort_dir")),
	}

	if filter.Tier == "" {
		filter.Tier = strings.TrimSpace(c.QueryParam("tier_eksposur"))
	}
	if filter.Product == "" {
		filter.Product = strings.TrimSpace(c.QueryParam("produk"))
	}
	if filter.Pengelola == "" {
		filter.Pengelola = strings.TrimSpace(c.QueryParam("officer_id"))
	}

	resp, err := h.usecase.GetAllCommitments(c.Request().Context(), filter)
	if err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	return RespondWithSuccess(c, http.StatusOK, "Commitment data retrieved successfully", resp)
}


// GetCommitmentHistory handles GET /api/v1/commitments/:account_no
func (h *CommitmentHandler) GetCommitmentHistory(c echo.Context) error {
	accountNo := c.Param("account_no")
	if accountNo == "" {
		return RespondWithError(c, http.StatusBadRequest, "account_no parameter is required")
	}

	history, err := h.usecase.GetCommitmentsByAccountNo(c.Request().Context(), accountNo)
	if err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	return RespondWithSuccess(c, http.StatusOK, "Commitment history retrieved successfully", history)
}

// UpdateCommitment handles PUT /api/v1/commitments/:id
func (h *CommitmentHandler) UpdateCommitment(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return RespondWithError(c, http.StatusBadRequest, "id parameter is required")
	}

	var req domain.UpdateCommitmentRequest
	if err := c.Bind(&req); err != nil {
		return RespondWithError(c, http.StatusBadRequest, "Invalid JSON request body")
	}

	commitment, err := h.usecase.UpdateCommitment(c.Request().Context(), id, req)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	return RespondWithSuccess(c, http.StatusOK, "Commitment updated successfully", commitment)
}

// ImportCommitmentsExcel handles POST /api/v1/commitments/import-excel
func (h *CommitmentHandler) ImportCommitmentsExcel(c echo.Context) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		fileHeader, err = c.FormFile("excel")
		if err != nil {
			return RespondWithError(c, http.StatusBadRequest, "excel file is required (form key 'file' or 'excel')")
		}
	}

	src, err := fileHeader.Open()
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, "failed to open uploaded file: "+err.Error())
	}
	defer src.Close()

	result, err := h.usecase.ImportCommitmentsFromExcel(c.Request().Context(), src)
	if err != nil {
		return RespondWithError(c, http.StatusBadRequest, err.Error())
	}

	return RespondWithSuccess(c, http.StatusOK, "Excel file processed successfully", result)
}

func (h *CommitmentHandler) getFilterFromQuery(c echo.Context) domain.CommitmentFilter {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	filter := domain.CommitmentFilter{
		Page:      page,
		Limit:     limit,
		AccountNo: strings.TrimSpace(c.QueryParam("account_no")),
		Status:    strings.TrimSpace(c.QueryParam("status")),
		Search:    strings.TrimSpace(c.QueryParam("search")),
		Tier:      strings.TrimSpace(c.QueryParam("tier")),
		Pengelola: strings.TrimSpace(c.QueryParam("pengelola")),
		Product:   strings.TrimSpace(c.QueryParam("product")),
		SortBy:    strings.TrimSpace(c.QueryParam("sort_by")),
		SortDir:   strings.TrimSpace(c.QueryParam("sort_dir")),
	}

	if filter.Tier == "" {
		filter.Tier = strings.TrimSpace(c.QueryParam("tier_eksposur"))
	}
	if filter.Product == "" {
		filter.Product = strings.TrimSpace(c.QueryParam("produk"))
	}
	if filter.Pengelola == "" {
		filter.Pengelola = strings.TrimSpace(c.QueryParam("officer_id"))
	}
	return filter
}

// ExportCommitmentsExcel handles GET /api/v1/commitments/export/excel
func (h *CommitmentHandler) ExportCommitmentsExcel(c echo.Context) error {
	filter := h.getFilterFromQuery(c)
	data, err := h.usecase.ExportCommitmentsExcel(c.Request().Context(), filter)
	if err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set(echo.HeaderContentDisposition, `attachment; filename="buma_commitments_report.xlsx"`)
	return c.Blob(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", data)
}

// ExportCommitmentsPDF handles GET /api/v1/commitments/export/pdf
func (h *CommitmentHandler) ExportCommitmentsPDF(c echo.Context) error {
	filter := h.getFilterFromQuery(c)
	data, err := h.usecase.ExportCommitmentsPDF(c.Request().Context(), filter)
	if err != nil {
		return RespondWithError(c, http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/pdf")
	c.Response().Header().Set(echo.HeaderContentDisposition, `attachment; filename="buma_commitments_report.pdf"`)
	return c.Blob(http.StatusOK, "application/pdf", data)
}




