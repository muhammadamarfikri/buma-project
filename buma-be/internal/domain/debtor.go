package domain

import "time"

// DebtorAccount represents a debtor facility item with officer assignment & latest commitment details
type DebtorAccount struct {
	AccountNo          string    `json:"noRekening"`
	DebtorID           string    `json:"id"`
	DebtorName         string    `json:"namaDebitur"`
	ProductType        string    `json:"produk"`
	CreditLimit        float64   `json:"limit"`
	OutstandingBalance float64   `json:"bakiDebet"`
	ExposureTier       string    `json:"tierEksposur"`
	OfficerID          string    `json:"officerId,omitempty"`
	OfficerName        string    `json:"pengelola"`
	PairingTeam        string    `json:"pairing"`
	CommitmentDate     string    `json:"tglKomitmen"`
	CommitmentStatus   string    `json:"statusKomitmen"`
	Remarks            string    `json:"keterangan"`
	NominalCommitment  float64   `json:"nominalKomitmen,omitempty"`
	Phone              string    `json:"phone"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

// DebtorFilter holds search and pagination criteria for GET /api/v1/monitoring
type DebtorFilter struct {
	Page        int    `json:"page"`
	Limit       int    `json:"limit"`
	Search      string `json:"search"`
	Tier        string `json:"tier"`
	OfficerID   string `json:"officer_id"`
	OfficerName string `json:"pengelola"`
	Product     string `json:"produk"`
	SortBy      string `json:"sort_by"`
	SortDir     string `json:"sort_dir"`
}

// MonitoringMetrics summarizes collection metrics for the dashboard cards
type MonitoringMetrics struct {
	TotalDebtors     int     `json:"totalDebtors"`
	TotalBakiDebet   float64 `json:"totalBakiDebet"`
	Tier1Count       int     `json:"tier1Count"`
	Tier1Ratio       float64 `json:"tier1Ratio"`
	CommitmentsToday int     `json:"commitmentsToday"`
}

// PaginationMeta contains pagination details
type PaginationMeta struct {
	CurrentPage  int `json:"currentPage"`
	ItemsPerPage int `json:"itemsPerPage"`
	TotalItems   int `json:"totalItems"`
	TotalPages   int `json:"totalPages"`
}

// MonitoringResponse is the standard response for GET /api/v1/monitoring
type MonitoringResponse struct {
	Metrics    MonitoringMetrics `json:"metrics"`
	Data       []DebtorAccount   `json:"data"`
	Pagination PaginationMeta    `json:"pagination"`
}
