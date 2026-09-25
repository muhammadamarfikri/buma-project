package usecase_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/xuri/excelize/v2"

	"buma-be/internal/repository/memory"
	"buma-be/internal/usecase"
)


func createTestExcel(rows [][]interface{}) (*bytes.Buffer, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	for rowIdx, row := range rows {
		for colIdx, val := range row {
			cell, _ := excelize.CoordinatesToCellName(colIdx+1, rowIdx+1)
			_ = f.SetCellValue(sheet, cell, val)
		}
	}

	buf := new(bytes.Buffer)
	if err := f.Write(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func TestImportCommitmentsFromExcel(t *testing.T) {
	commitmentRepo := memory.NewInMemoryCommitmentRepository()
	debtorRepo := memory.NewInMemoryDebtorRepository()
	officerRepo := memory.NewInMemoryOfficerRepository()

	uc := usecase.NewCommitmentUsecase(commitmentRepo, debtorRepo, officerRepo)
	ctx := context.Background()

	t.Run("Valid Excel with Header Row", func(t *testing.T) {
		rows := [][]interface{}{
			{"Account No", "Debtor Name", "Officer ID", "Officer Pair ID", "Product", "Exposure Tier", "Outstanding Balance", "Credit Limit", "Commitment Date", "Status", "Reason", "Remarks"},
			{"1029384756", "PT Sukses Maju", "OFF-001", "OFF-002", "Kredit Modal Kerja", "Tier 1", "4500000000", "5000000000", "2026-09-20", "Janji Bayar (PTP)", "Pelunasan Sebagian", "Excel Import Batch 1"},
			{"5678901234", "CV Karya Bersama", "OFF-003", "OFF-004", "Kredit Investasi", "Tier 2", "1200000000", "2000000000", "2026-09-22", "Restrukturisasi", "Reschedule Angsuran", "Excel Import Batch 2"},
		}

		buf, err := createTestExcel(rows)
		if err != nil {
			t.Fatalf("Failed to create test excel: %v", err)
		}

		res, err := uc.ImportCommitmentsFromExcel(ctx, buf)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if res.TotalProcessed != 2 {
			t.Errorf("Expected TotalProcessed 2, got %d", res.TotalProcessed)
		}

		if res.TotalSuccess != 2 {
			t.Errorf("Expected TotalSuccess 2, got %d", res.TotalSuccess)
		}

		if len(res.Errors) != 0 {
			t.Errorf("Expected 0 errors, got %d", len(res.Errors))
		}
	})

	t.Run("Excel with Invalid and Missing Rows", func(t *testing.T) {
		rows := [][]interface{}{
			{"account_no", "debtor_name", "reason"},
			{"", "Missing Account", "Pelunasan"},
			{"1029384756", "Valid Account", ""},
		}

		buf, err := createTestExcel(rows)
		if err != nil {
			t.Fatalf("Failed to create test excel: %v", err)
		}

		res, err := uc.ImportCommitmentsFromExcel(ctx, buf)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if res.TotalProcessed != 2 {
			t.Errorf("Expected TotalProcessed 2, got %d", res.TotalProcessed)
		}

		if res.TotalFailed != 1 {
			t.Errorf("Expected TotalFailed 1, got %d", res.TotalFailed)
		}

		if res.TotalSuccess != 1 {
			t.Errorf("Expected TotalSuccess 1, got %d", res.TotalSuccess)
		}

		if len(res.Errors) != 1 {
			t.Errorf("Expected 1 row error, got %d", len(res.Errors))
		}
	})

	t.Run("Excel Import Create or Update Matching AccountNo and Reason", func(t *testing.T) {
		// First import to create initial commitment
		rows1 := [][]interface{}{
			{"Account No", "Debtor Name", "Officer ID", "Officer Pair ID", "Product", "Exposure Tier", "Outstanding Balance", "Credit Limit", "Commitment Date", "Status", "Reason", "Remarks"},
			{"1029384756", "PT Sukses Maju", "OFF-001", "OFF-002", "KMK", "Tier 1", "4000000000", "5000000000", "2026-09-20", "On Progress/Implementasi", "Pelunasan Sebagian", "Initial import"},
		}

		buf1, err := createTestExcel(rows1)
		if err != nil {
			t.Fatalf("Failed to create test excel: %v", err)
		}

		res1, err := uc.ImportCommitmentsFromExcel(ctx, buf1)
		if err != nil || res1.TotalSuccess != 1 {
			t.Fatalf("Expected 1 success on initial import, got err=%v, res=%+v", err, res1)
		}
		origID := res1.Successes[0].ID

		// Second import with SAME account_no AND SAME reason -> should UPDATE existing commitment
		rows2 := [][]interface{}{
			{"Account No", "Debtor Name", "Officer ID", "Officer Pair ID", "Product", "Exposure Tier", "Outstanding Balance", "Credit Limit", "Commitment Date", "Status", "Reason", "Remarks"},
			{"1029384756", "PT Sukses Maju Updated", "OFF-001", "OFF-002", "KMK", "Tier 1", "4500000000", "5000000000", "2026-09-25", "Telah Dilaksanakan", "Pelunasan Sebagian", "Updated remarks via re-import"},
		}

		buf2, err := createTestExcel(rows2)
		if err != nil {
			t.Fatalf("Failed to create test excel: %v", err)
		}

		res2, err := uc.ImportCommitmentsFromExcel(ctx, buf2)
		if err != nil || res2.TotalSuccess != 1 {
			t.Fatalf("Expected 1 success on update import, got err=%v, res=%+v", err, res2)
		}

		updatedItem := res2.Successes[0]
		if updatedItem.ID != origID {
			t.Errorf("Expected updated commitment to keep same ID %s, got %s", origID, updatedItem.ID)
		}
		if updatedItem.Remarks != "Updated remarks via re-import" {
			t.Errorf("Expected updated remarks, got %s", updatedItem.Remarks)
		}
		if updatedItem.CommitmentDate != "2026-09-25" {
			t.Errorf("Expected updated date 2026-09-25, got %s", updatedItem.CommitmentDate)
		}

		// Third import with SAME account_no but DIFFERENT reason -> should CREATE new commitment
		rows3 := [][]interface{}{
			{"Account No", "Debtor Name", "Officer ID", "Officer Pair ID", "Product", "Exposure Tier", "Outstanding Balance", "Credit Limit", "Commitment Date", "Status", "Reason", "Remarks"},
			{"1029384756", "PT Sukses Maju", "OFF-001", "OFF-002", "KMK", "Tier 1", "4500000000", "5000000000", "2026-10-01", "On Progress/Implementasi", "Restrukturisasi Baru", "Different reason creates new record"},
		}

		buf3, err := createTestExcel(rows3)
		if err != nil {
			t.Fatalf("Failed to create test excel: %v", err)
		}

		res3, err := uc.ImportCommitmentsFromExcel(ctx, buf3)
		if err != nil || res3.TotalSuccess != 1 {
			t.Fatalf("Expected 1 success on new reason import, got err=%v, res=%+v", err, res3)
		}

		newItem := res3.Successes[0]
		if newItem.ID == origID {
			t.Errorf("Expected new commitment to have different ID, but got same ID %s", origID)
		}
	})
}
