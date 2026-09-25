package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"

	"buma-be/internal/domain"
	"buma-be/internal/repository"
)

type CommitmentUsecase interface {
	CreateCommitment(ctx context.Context, req domain.CreateCommitmentRequest) (*domain.CollectionCommitment, error)
	GetCommitmentsByAccountNo(ctx context.Context, accountNo string) ([]domain.CollectionCommitment, error)
	GetAllCommitments(ctx context.Context, filter domain.CommitmentFilter) (*domain.PaginatedCommitmentResponse, error)
	UpdateCommitment(ctx context.Context, id string, req domain.UpdateCommitmentRequest) (*domain.CollectionCommitment, error)
	ImportCommitmentsFromExcel(ctx context.Context, reader io.Reader) (*domain.ExcelImportResult, error)
}

type commitmentUsecase struct {
	commitmentRepo repository.CommitmentRepository
	debtorRepo     repository.DebtorRepository
	officerRepo    repository.OfficerRepository
}

func NewCommitmentUsecase(commitmentRepo repository.CommitmentRepository, debtorRepo repository.DebtorRepository, officerRepo repository.OfficerRepository) CommitmentUsecase {
	return &commitmentUsecase{
		commitmentRepo: commitmentRepo,
		debtorRepo:     debtorRepo,
		officerRepo:    officerRepo,
	}
}

func (u *commitmentUsecase) CreateCommitment(ctx context.Context, req domain.CreateCommitmentRequest) (*domain.CollectionCommitment, error) {

	// Check the officer exists
	if err := u.isOfficerExist(ctx, req); err != nil {
		return nil, err
	}

	// Fetch existing commitments for this account to check if reason exists
	commitments, err := u.commitmentRepo.GetCommitmentsByAccountNo(ctx, req.AccountNo)
	if err != nil {
		return nil, err
	}

	var res *domain.CollectionCommitment
	// If reason exists, update existing record. Otherwise, create a new record.
	if u.isCommitmentReasonExist(commitments, req.Reason) {
		res, err = u.commitmentRepo.UpdateCommitmentByReason(ctx, req)
	} else {
		res, err = u.commitmentRepo.CreateCommitment(ctx, req)
	}

	if err != nil || res == nil {
		return nil, err
	}

	enriched := u.enrichOfficerDetails(ctx, []domain.CollectionCommitment{*res})
	return &enriched[0], nil
}

func (u *commitmentUsecase) GetCommitmentsByAccountNo(ctx context.Context, accountNo string) ([]domain.CollectionCommitment, error) {
	if accountNo == "" {
		return nil, errors.New("account_no is required")
	}
	list, err := u.commitmentRepo.GetCommitmentsByAccountNo(ctx, accountNo)
	if err != nil {
		return nil, err
	}
	return u.enrichOfficerDetails(ctx, list), nil
}

func (u *commitmentUsecase) GetAllCommitments(ctx context.Context, filter domain.CommitmentFilter) (*domain.PaginatedCommitmentResponse, error) {
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 {
		filter.Limit = 10
	}
	resp, err := u.commitmentRepo.GetCommitments(ctx, filter)
	if err != nil || resp == nil {
		return nil, err
	}
	resp.Data = u.enrichOfficerDetails(ctx, resp.Data)
	return resp, nil
}

func (u *commitmentUsecase) enrichOfficerDetails(ctx context.Context, list []domain.CollectionCommitment) []domain.CollectionCommitment {
	if u.officerRepo == nil || len(list) == 0 {
		return list
	}

	officers, err := u.officerRepo.GetAllOfficers(ctx)
	if err != nil || len(officers) == 0 {
		return list
	}

	officerMap := make(map[string]domain.Officer)
	for _, o := range officers {
		officerMap[o.OfficerID] = o
	}

	for i := range list {
		if list[i].OfficerID != "" {
			if o, ok := officerMap[list[i].OfficerID]; ok {
				list[i].OfficerName = o.FullName
			} else if list[i].OfficerName == "" {
				list[i].OfficerName = list[i].OfficerID
			}
		}
		if list[i].OfficerPairID != "" {
			if o, ok := officerMap[list[i].OfficerPairID]; ok {
				list[i].OfficerPairName = o.FullName
			} else if list[i].OfficerPairName == "" {
				list[i].OfficerPairName = list[i].OfficerPairID
			}
		}
	}

	return list
}


func (u *commitmentUsecase) UpdateCommitment(ctx context.Context, id string, req domain.UpdateCommitmentRequest) (*domain.CollectionCommitment, error) {
	if id == "" {
		return nil, errors.New("commitment id is required")
	}

	if u.officerRepo != nil {
		if req.OfficerID != "" {
			officer, err := u.officerRepo.GetOfficerByID(ctx, req.OfficerID)
			if err != nil {
				return nil, err
			}
			if officer == nil {
				return nil, errors.New("officer not found")
			}
		}

		if req.OfficerPairID != "" {
			officerPair, err := u.officerRepo.GetOfficerByID(ctx, req.OfficerPairID)
			if err != nil {
				return nil, err
			}
			if officerPair == nil {
				return nil, errors.New("officer pair not found")
			}
		}
	}

	res, err := u.commitmentRepo.UpdateCommitmentByID(ctx, id, req)
	if err != nil || res == nil {
		return nil, err
	}

	enriched := u.enrichOfficerDetails(ctx, []domain.CollectionCommitment{*res})
	return &enriched[0], nil
}

func (u *commitmentUsecase) ImportCommitmentsFromExcel(ctx context.Context, reader io.Reader) (*domain.ExcelImportResult, error) {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to open excel file: %w", err)
	}
	defer f.Close()

	sheetList := f.GetSheetList()
	if len(sheetList) == 0 {
		return nil, errors.New("excel file contains no sheets")
	}

	result := &domain.ExcelImportResult{
		Errors:    []domain.ExcelImportRowError{},
		Successes: []domain.CollectionCommitment{},
	}

	totalRowIndex := 0

	for _, sheetName := range sheetList {
		rows, err := f.GetRows(sheetName)
		if err != nil || len(rows) == 0 {
			continue
		}

		headerMap := make(map[string]int)
		headerRowIdx := -1

		// Search for header row in the first 10 rows
		for rowIdx := 0; rowIdx < len(rows) && rowIdx < 10; rowIdx++ {
			row := rows[rowIdx]
			tempMap := make(map[string]int)
			for colIdx, colVal := range row {
				colClean := normalizeHeader(colVal)
				if colClean != "" {
					tempMap[colClean] = colIdx
				}
			}
			if containsAnyHeader(tempMap, "accountno", "account_no", "accountnumber", "norekening", "norek", "no_rek", "rekening", "account", "cif") {
				headerMap = tempMap
				headerRowIdx = rowIdx
				break
			}
		}

		startRowIdx := 1
		if headerRowIdx >= 0 {
			startRowIdx = headerRowIdx + 1
		} else {
			// Default column mapping if no header row found
			startRowIdx = 0
			headerMap = map[string]int{
				"accountno":          0,
				"debtorname":         1,
				"officerid":          2,
				"officerpairid":      3,
				"product":            4,
				"exposuretier":       5,
				"outstandingbalance": 6,
				"creditlimit":        7,
				"commitmentdate":     8,
				"status":             9,
				"reason":             10,
				"remarks":            11,
			}
		}

		for i := startRowIdx; i < len(rows); i++ {
			row := rows[i]
			totalRowIndex++
			if isRowEmpty(row) {
				continue
			}

			accountNo := getColValue(row, headerMap, "accountno", "account_no", "accountnumber", "norekening", "norek", "no_rek", "rekening", "account", "cif")
			if accountNo == "CIF" || accountNo == "No Rek" || accountNo == "No. Rek" || accountNo == "No Rekening" {
				continue
			}

			result.TotalProcessed++

			if accountNo == "" {
				result.TotalFailed++
				result.Errors = append(result.Errors, domain.ExcelImportRowError{
					Row:   totalRowIndex,
					Error: "account_no is required",
				})
				continue
			}

			debtorName := getColValue(row, headerMap, "debtorname", "debtor_name", "namadebitur", "nama_debitur", "debtor")
			officerID := getColValue(row, headerMap, "officerid", "officer_id", "officer", "pengelola")
			officerPairID := getColValue(row, headerMap, "officerpairid", "officer_pair_id", "officerpair", "pairing")
			product := getColValue(row, headerMap, "product", "produk")
			exposureTier := getColValue(row, headerMap, "exposuretier", "exposure_tier", "tiereksposur", "tier")
			outstandingBalance := parseFloat(getColValue(row, headerMap, "outstandingbalance", "outstanding_balance", "outstanding", "bakidebet", "baki_debet", "bade"))
			creditLimit := parseFloat(getColValue(row, headerMap, "creditlimit", "credit_limit", "limit"))
			commitmentDateRaw := getColValue(row, headerMap, "commitmentdate", "commitment_date", "tglkomitmen", "tanggalkomitmen", "date")
			commitmentDate := parseExcelDate(commitmentDateRaw)

			reason := getColValue(row, headerMap, "reason", "alasan")
			if reason == "-" {
				reason = ""
			}
			status := getColValue(row, headerMap, "status", "keterangan")
			if status == "-" {
				status = ""
			}
			remarks := getColValue(row, headerMap, "remarks", "catatan", "staging")
			if remarks == "-" {
				remarks = ""
			}

			req := domain.CreateCommitmentRequest{
				AccountNo:          accountNo,
				DebtorName:         debtorName,
				OfficerID:          officerID,
				OfficerPairID:      officerPairID,
				Product:            product,
				ExposureTier:       exposureTier,
				OutstandingBalance: outstandingBalance,
				CreditLimit:        creditLimit,
				CommitmentDate:     commitmentDate,
				Status:             status,
				Reason:             reason,
				Remarks:            remarks,
			}

			commitment, err := u.CreateCommitment(ctx, req)
			if err != nil {
				result.TotalFailed++
				result.Errors = append(result.Errors, domain.ExcelImportRowError{
					Row:   totalRowIndex,
					Error: err.Error(),
				})
			} else {
				result.TotalSuccess++
				result.Successes = append(result.Successes, *commitment)
			}
		}
	}

	if result.TotalProcessed == 0 {
		return nil, errors.New("excel file contains no valid data rows")
	}

	return result, nil
}

func normalizeHeader(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "_", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, ".", "")
	return strings.TrimSpace(s)
}

func containsAnyHeader(headerMap map[string]int, keys ...string) bool {
	for _, k := range keys {
		if _, ok := headerMap[normalizeHeader(k)]; ok {
			return true
		}
	}
	return false
}

func getColValue(row []string, headerMap map[string]int, keys ...string) string {
	for _, k := range keys {
		normKey := normalizeHeader(k)
		if colIdx, ok := headerMap[normKey]; ok {
			if colIdx < len(row) {
				return strings.TrimSpace(row[colIdx])
			}
		}
	}
	return ""
}

func isRowEmpty(row []string) bool {
	for _, cell := range row {
		if strings.TrimSpace(cell) != "" {
			return false
		}
	}
	return true
}

func parseFloat(val string) float64 {
	val = strings.TrimSpace(val)
	val = strings.ReplaceAll(val, ",", "")
	if val == "" {
		return 0
	}
	f, _ := strconv.ParseFloat(val, 64)
	return f
}

func parseExcelDate(val string) string {
	val = strings.TrimSpace(val)
	if val == "" {
		return ""
	}

	// Try YYYY-MM-DD
	if t, err := time.Parse("2006-01-02", val); err == nil {
		return t.Format("2006-01-02")
	}
	// Try MM-DD-YY (e.g. 09-11-26 -> 2026-09-11)
	if t, err := time.Parse("01-02-06", val); err == nil {
		return t.Format("2006-01-02")
	}
	// Try MM/DD/YY
	if t, err := time.Parse("01/02/06", val); err == nil {
		return t.Format("2006-01-02")
	}
	// Try MM/DD/YYYY
	if t, err := time.Parse("01/02/2006", val); err == nil {
		return t.Format("2006-01-02")
	}
	// Try DD-MM-YYYY
	if t, err := time.Parse("02-01-2006", val); err == nil {
		return t.Format("2006-01-02")
	}
	// Try Excel serial date number
	if num, err := strconv.ParseFloat(val, 64); err == nil && num > 30000 && num < 100000 {
		excelEpoch := time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
		t := excelEpoch.Add(time.Duration(num * float64(24*time.Hour)))
		return t.Format("2006-01-02")
	}

	return ""
}

func (u *commitmentUsecase) isCommitmentReasonExist(data []domain.CollectionCommitment, reason string) bool {
	reasonTrimmed := strings.TrimSpace(reason)
	for _, v := range data {
		if strings.EqualFold(strings.TrimSpace(v.Reason), reasonTrimmed) {
			return true
		}
	}
	return false
}

func (u *commitmentUsecase) isOfficerExist(ctx context.Context, req domain.CreateCommitmentRequest) error {
	return nil
}


