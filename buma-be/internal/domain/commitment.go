package domain

import "time"

type CollectionCommitment struct {
	ID                 string    `json:"id"`
	AccountNo          string    `json:"account_no"`
	DebtorName         string    `json:"debtor_name,omitempty"`
	CommitmentDate     string    `json:"commitment_date"`
	OfficerID          string    `json:"officer_id"`
	OfficerPairID      string    `json:"officer_pair_id"`
	OfficerName        string    `json:"officer_name,omitempty"`
	OfficerPairName    string    `json:"officer_pair_name,omitempty"`
	Status             string    `json:"status"`
	Remarks            string    `json:"remarks"`
	Nominal            float64   `json:"nominal"`
	Reason             string    `json:"reason"`
	CreditLimit        float64   `json:"credit_limit,omitempty"`
	OutstandingBalance float64   `json:"outstanding_balance,omitempty"`
	ExposureTier       string    `json:"exposure_tier,omitempty"`
	Product            string    `json:"product,omitempty"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type CreateCommitmentRequest struct {
	AccountNo          string  `json:"account_no" validate:"required"`
	DebtorName         string  `json:"debtor_name"`
	OfficerID          string  `json:"officer_id"`
	OfficerPairID      string  `json:"officer_pair_id"`
	Product            string  `json:"product"`
	ExposureTier       string  `json:"exposure_tier"`
	OutstandingBalance float64 `json:"outstanding_balance,omitempty"`
	CreditLimit        float64 `json:"credit_limit,omitempty"`
	CommitmentDate     string  `json:"commitment_date"`
	Status             string  `json:"status"`
	Reason             string  `json:"reason"`
	Remarks            string  `json:"remarks"`
}

type UpdateCommitmentRequest struct {
	AccountNo          string  `json:"account_no,omitempty"`
	Product            string  `json:"product,omitempty"`
	CommitmentDate     string  `json:"commitment_date"`
	Status             string  `json:"status"`
	Reason             string  `json:"reason"`
	Remarks            string  `json:"remarks"`
	Nominal            float64 `json:"nominal"`
	OfficerID          string  `json:"officer_id"`
	OfficerPairID      string  `json:"officer_pair_id"`
	CreditLimit        float64 `json:"credit_limit"`
	OutstandingBalance float64 `json:"outstanding_balance"`
	ExposureTier       string  `json:"exposure_tier"`
}

type CommitmentResponse struct {
	Message    string               `json:"message"`
	Commitment CollectionCommitment `json:"commitment"`
}

type CommitmentFilter struct {
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	AccountNo string `json:"account_no"`
	Status    string `json:"status"`
	Search    string `json:"search"`
	Tier      string `json:"tier"`
	Pengelola string `json:"pengelola"`
	Product   string `json:"product"`
	SortBy    string `json:"sort_by"`
	SortDir   string `json:"sort_dir"`
}

type PaginatedCommitmentResponse struct {
	Data       []CollectionCommitment `json:"data"`
	Pagination PaginationMeta         `json:"pagination"`
}

type ExcelImportRowError struct {
	Row   int    `json:"row"`
	Error string `json:"error"`
}

type ExcelImportResult struct {
	TotalProcessed int                    `json:"total_processed"`
	TotalSuccess   int                    `json:"total_success"`
	TotalFailed    int                    `json:"total_failed"`
	Errors         []ExcelImportRowError  `json:"errors,omitempty"`
	Successes      []CollectionCommitment `json:"successes,omitempty"`
}
