package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"buma-be/internal/handler"
	"buma-be/internal/validator"
)

type RouterDependencies struct {
	MonitoringHandler *handler.MonitoringHandler
	CommitmentHandler *handler.CommitmentHandler
	WordingHandler    *handler.WordingHandler
	OfficerHandler    *handler.OfficerHandler
}

func SetupRouter(deps RouterDependencies) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// Register Struct Validator
	e.Validator = validator.NewCustomValidator()

	// 1. Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 2. CORS Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderXRequestedWith, "X-CSRF-Token"},
	}))

	// 3. Health Check
	e.GET("/health", func(c echo.Context) error {
		return handler.RespondWithSuccess(c, http.StatusOK, "BUMA Debt Collection Backend Service is healthy", map[string]string{
			"status":  "UP",
			"service": "buma-be",
		})
	})

	// 4. API v1 Routes
	v1 := e.Group("/api/v1")
	{
		// Monitoring Endpoints
		v1.GET("/monitoring", deps.MonitoringHandler.GetMonitoring)
		v1.GET("/monitoring/:account_no", deps.MonitoringHandler.GetMonitoringByAccountNo)

		// Commitment Endpoints
		v1.POST("/commitments", deps.CommitmentHandler.CreateCommitment)
		v1.POST("/commitments/import-excel", deps.CommitmentHandler.ImportCommitmentsExcel)
		v1.GET("/commitments/export/excel", deps.CommitmentHandler.ExportCommitmentsExcel)
		v1.GET("/commitments/export/pdf", deps.CommitmentHandler.ExportCommitmentsPDF)
		v1.GET("/commitments", deps.CommitmentHandler.GetAllCommitments)
		v1.GET("/commitments/:account_no", deps.CommitmentHandler.GetCommitmentHistory)
		v1.PUT("/commitments/:id", deps.CommitmentHandler.UpdateCommitment)


		// Wording Generator Endpoint
		v1.GET("/templates/wording", deps.WordingHandler.GenerateWording)

		// Officer Management Endpoints
		if deps.OfficerHandler != nil {
			v1.POST("/officers", deps.OfficerHandler.CreateOfficer)
			v1.GET("/officers", deps.OfficerHandler.GetAllOfficers)
			v1.GET("/officers/:officer_id", deps.OfficerHandler.GetOfficerByID)
		}
	}

	return e
}

