package wording

import (
	"strings"
	"testing"

	"buma-be/internal/domain"
)

func TestGenerateCollectionWording(t *testing.T) {
	sampleDebtor := domain.DebtorAccount{
		AccountNo:          "1029384756",
		DebtorID:           "DBT-001",
		DebtorName:         "PT Nusantara Jaya Abadi",
		ProductType:        "KMK",
		CreditLimit:        5000000000,
		OutstandingBalance: 4250000000,
		ExposureTier:       "Tier 1",
		OfficerName:        "Budi Santoso",
		PairingTeam:        "Desk 01",
		CommitmentDate:     "2026-09-06",
		CommitmentStatus:   "Janji Bayar (PTP)",
		Phone:              "081298765432",
	}

	t.Run("WhatsApp Soft Tone", func(t *testing.T) {
		res := GenerateCollectionWording(sampleDebtor, "whatsapp", "soft")
		if !strings.Contains(res, "PT Nusantara Jaya Abadi") {
			t.Errorf("Expected debtor name in wording, got: %s", res)
		}
		if !strings.Contains(res, "Rp 4.250.000.000") {
			t.Errorf("Expected formatted balance in wording, got: %s", res)
		}
	})

	t.Run("WhatsApp Urgent Tone", func(t *testing.T) {
		res := GenerateCollectionWording(sampleDebtor, "whatsapp", "urgent")
		if !strings.Contains(res, "SURAT PERINGATAN / SOMASI PRA-HUKUM") {
			t.Errorf("Expected urgent header, got: %s", res)
		}
	})

	t.Run("Email Template", func(t *testing.T) {
		res := GenerateCollectionWording(sampleDebtor, "email", "firm")
		if !strings.Contains(res, "Subjek: [BUMA COLLECTION]") {
			t.Errorf("Expected email subject, got: %s", res)
		}
	})

	t.Run("FormatIDR Utility", func(t *testing.T) {
		formatted := FormatIDR(1250000000)
		if formatted != "Rp 1.250.000.000" {
			t.Errorf("Expected Rp 1.250.000.000, got: %s", formatted)
		}
	})
}
