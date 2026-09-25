package handler_test

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/xuri/excelize/v2"

	"buma-be/internal/domain"
	"buma-be/internal/handler"
	"buma-be/internal/repository/memory"
	"buma-be/internal/router"
	"buma-be/internal/usecase"
)


func setupTestServer() http.Handler {
	debtorRepo := memory.NewInMemoryDebtorRepository()
	officerRepo := memory.NewInMemoryOfficerRepository()
	commitmentRepo := memory.NewInMemoryCommitmentRepository()

	monitoringUc := usecase.NewMonitoringUsecase(debtorRepo)
	commitmentUc := usecase.NewCommitmentUsecase(commitmentRepo, debtorRepo, officerRepo)
	wordingUc := usecase.NewWordingUsecase(debtorRepo)
	officerUc := usecase.NewOfficerUsecase(officerRepo)

	monitoringHnd := handler.NewMonitoringHandler(monitoringUc)
	commitmentHnd := handler.NewCommitmentHandler(commitmentUc)
	wordingHnd := handler.NewWordingHandler(wordingUc)
	officerHnd := handler.NewOfficerHandler(officerUc)

	deps := router.RouterDependencies{
		MonitoringHandler: monitoringHnd,
		CommitmentHandler: commitmentHnd,
		WordingHandler:    wordingHnd,
		OfficerHandler:    officerHnd,
	}

	return router.SetupRouter(deps)
}

func TestHealthCheck(t *testing.T) {
	ts := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	ts.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", rec.Code)
	}

	var res handler.APIResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if !res.Success {
		t.Fatalf("Expected success true, got false")
	}
}

func TestGetMonitoringEndpoint(t *testing.T) {
	ts := setupTestServer()

	t.Run("GET /api/v1/monitoring", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/monitoring?page=1&limit=5", nil)
		rec := httptest.NewRecorder()

		ts.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("Expected status 200 OK, got %d", rec.Code)
		}

		var resp domain.MonitoringResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
			t.Fatalf("Failed to unmarshal monitoring response: %v", err)
		}

		if resp.Metrics.TotalDebtors != 12 {
			t.Errorf("Expected 12 total debtors in metrics, got %d", resp.Metrics.TotalDebtors)
		}

		if len(resp.Data) != 5 {
			t.Errorf("Expected 5 items in paginated page 1, got %d", len(resp.Data))
		}
	})

	t.Run("GET /api/v1/monitoring/1029384756", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/monitoring/1029384756", nil)
		rec := httptest.NewRecorder()

		ts.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("Expected status 200 OK, got %d", rec.Code)
		}

		var res handler.APIResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if !res.Success {
			t.Fatalf("Expected success true, got false")
		}
	})
}

func TestGenerateWordingEndpoint(t *testing.T) {
	ts := setupTestServer()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/templates/wording?account_no=1029384756&channel=whatsapp&tone=urgent", nil)
	rec := httptest.NewRecorder()

	ts.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", rec.Code)
	}

	var res handler.APIResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	dataMap, ok := res.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("Expected map[string]interface{} data response")
	}

	wordingStr, _ := dataMap["wording"].(string)
	if wordingStr == "" {
		t.Errorf("Expected wording string in response, got empty")
	}
}

func TestCommitmentEndpointValidation(t *testing.T) {
	ts := setupTestServer()

	body := []byte(`{}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/commitments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	ts.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("Expected status 400 Bad Request for missing account_no, got %d", rec.Code)
	}
}

func TestOfficerEndpoints(t *testing.T) {
	ts := setupTestServer()

	t.Run("POST /api/v1/officers - Add New Officer", func(t *testing.T) {
		body := []byte(`{
			"officer_id": "OFF-005",
			"full_name": "Rizky Ramadhan",
			"email": "rizky.ramadhan@buma.co.id",
			"pairing_team": "Desk 05",
			"phone": "081299990000"
		}`)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/officers", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		ts.ServeHTTP(rec, req)

		if rec.Code != http.StatusCreated {
			t.Fatalf("Expected status 201 Created, got %d: %s", rec.Code, rec.Body.String())
		}

		var res handler.APIResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if !res.Success {
			t.Fatalf("Expected success true, got false")
		}
	})

	t.Run("GET /api/v1/officers - Get All Officers", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/officers", nil)
		rec := httptest.NewRecorder()

		ts.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("Expected status 200 OK, got %d", rec.Code)
		}
	})
}

func TestGetAllCommitmentsEndpoint(t *testing.T) {
	ts := setupTestServer()

	t.Run("GET /api/v1/commitments - Get All Commitments Paginated", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/commitments?page=1&limit=2", nil)
		rec := httptest.NewRecorder()

		ts.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("Expected status 200 OK, got %d", rec.Code)
		}

		var res handler.APIResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if !res.Success {
			t.Fatalf("Expected success true, got false")
		}

		dataBytes, err := json.Marshal(res.Data)
		if err != nil {
			t.Fatalf("Failed to marshal data field: %v", err)
		}

		var paginatedResp domain.PaginatedCommitmentResponse
		if err := json.Unmarshal(dataBytes, &paginatedResp); err != nil {
			t.Fatalf("Failed to unmarshal paginated response: %v", err)
		}

		if len(paginatedResp.Data) != 2 {
			t.Errorf("Expected 2 items in page 1 limit 2, got %d", len(paginatedResp.Data))
		}

		if paginatedResp.Pagination.CurrentPage != 1 {
			t.Errorf("Expected CurrentPage 1, got %d", paginatedResp.Pagination.CurrentPage)
		}

		if paginatedResp.Pagination.ItemsPerPage != 2 {
			t.Errorf("Expected ItemsPerPage 2, got %d", paginatedResp.Pagination.ItemsPerPage)
		}
	})
}

func TestUpdateCommitmentEndpoint(t *testing.T) {
	ts := setupTestServer()

	t.Run("PUT /api/v1/commitments/:id - Update Commitment Success", func(t *testing.T) {
		createBody := []byte(`{
			"account_no": "1029384756",
			"commitment_date": "2026-09-08",
			"status": "Janji Bayar (PTP)",
			"reason": "Restrukturisasi",
			"remarks": "Initial remarks",
			"nominal": 100000000
		}`)
		reqCreate := httptest.NewRequest(http.MethodPost, "/api/v1/commitments", bytes.NewBuffer(createBody))
		reqCreate.Header.Set("Content-Type", "application/json")
		recCreate := httptest.NewRecorder()

		ts.ServeHTTP(recCreate, reqCreate)
		if recCreate.Code != http.StatusCreated {
			t.Fatalf("Expected status 201 Created, got %d: %s", recCreate.Code, recCreate.Body.String())
		}

		var resCreate handler.APIResponse
		if err := json.Unmarshal(recCreate.Body.Bytes(), &resCreate); err != nil {
			t.Fatalf("Failed to decode create response: %v", err)
		}

		dataMap, ok := resCreate.Data.(map[string]interface{})
		if !ok {
			t.Fatalf("Expected map[string]interface{} in resCreate.Data")
		}
		createdID, _ := dataMap["id"].(string)

		updateBody := []byte(`{
			"commitment_date": "2026-09-15",
			"status": "Kept",
			"reason": "Restrukturisasi",
			"remarks": "Updated remarks post committee",
			"nominal": 120000000,
			"officer_id": "OFF-002",
			"officer_pair_id": "OFF-003",
			"credit_limit": 6000000000,
			"outstanding_balance": 4250000000,
			"exposure_tier": "Tier 1"
		}`)
		reqUpdate := httptest.NewRequest(http.MethodPut, "/api/v1/commitments/"+createdID, bytes.NewBuffer(updateBody))
		reqUpdate.Header.Set("Content-Type", "application/json")
		recUpdate := httptest.NewRecorder()

		ts.ServeHTTP(recUpdate, reqUpdate)
		if recUpdate.Code != http.StatusOK {
			t.Fatalf("Expected status 200 OK, got %d: %s", recUpdate.Code, recUpdate.Body.String())
		}

		var resUpdate handler.APIResponse
		if err := json.Unmarshal(recUpdate.Body.Bytes(), &resUpdate); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if !resUpdate.Success {
			t.Fatalf("Expected success true, got false")
		}
	})
}

func TestImportCommitmentsExcelEndpoint(t *testing.T) {
	ts := setupTestServer()

	// Generate Excel in memory
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	_ = f.SetCellValue(sheet, "A1", "account_no")
	_ = f.SetCellValue(sheet, "B1", "reason")
	_ = f.SetCellValue(sheet, "C1", "remarks")
	_ = f.SetCellValue(sheet, "A2", "1029384756")
	_ = f.SetCellValue(sheet, "B2", "Restrukturisasi Excel Test")
	_ = f.SetCellValue(sheet, "C2", "Imported via HTTP Handler Test")

	excelBuf := new(bytes.Buffer)
	if err := f.Write(excelBuf); err != nil {
		t.Fatalf("Failed to write excel file: %v", err)
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "commitments.xlsx")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	_, _ = part.Write(excelBuf.Bytes())
	_ = writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/commitments/import-excel", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()

	ts.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d: %s", rec.Code, rec.Body.String())
	}

	var res handler.APIResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if !res.Success {
		t.Fatalf("Expected success true, got false")
	}
}

func TestGetAllCommitmentsWithFilters(t *testing.T) {
	ts := setupTestServer()

	t.Run("GET /api/v1/commitments with Tier and Product filter", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/commitments?tier=Tier%201&produk=KMK&page=1&limit=10", nil)
		rec := httptest.NewRecorder()

		ts.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Fatalf("Expected status 200 OK, got %d", rec.Code)
		}

		var res handler.APIResponse
		if err := json.Unmarshal(rec.Body.Bytes(), &res); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		if !res.Success {
			t.Fatalf("Expected success true, got false")
		}
	})
}



