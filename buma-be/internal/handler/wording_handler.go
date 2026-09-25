package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"buma-be/internal/usecase"
)

type WordingHandler struct {
	usecase usecase.WordingUsecase
}

func NewWordingHandler(u usecase.WordingUsecase) *WordingHandler {
	return &WordingHandler{usecase: u}
}

// GenerateWording handles GET /api/v1/templates/wording?account_no=xxx&channel=whatsapp&tone=firm
func (h *WordingHandler) GenerateWording(c echo.Context) error {
	accountNo := strings.TrimSpace(c.QueryParam("account_no"))
	channel := strings.TrimSpace(c.QueryParam("channel"))
	tone := strings.TrimSpace(c.QueryParam("tone"))

	if accountNo == "" {
		return RespondWithError(c, http.StatusBadRequest, "account_no query parameter is required")
	}

	wordingText, debtor, err := h.usecase.GenerateWording(c.Request().Context(), accountNo, channel, tone)
	if err != nil {
		return RespondWithError(c, http.StatusNotFound, err.Error())
	}

	payload := map[string]interface{}{
		"account_no": accountNo,
		"debtor":     debtor,
		"channel":    channel,
		"tone":       tone,
		"wording":    wordingText,
	}

	return RespondWithSuccess(c, http.StatusOK, "Collection wording draft generated successfully", payload)
}
