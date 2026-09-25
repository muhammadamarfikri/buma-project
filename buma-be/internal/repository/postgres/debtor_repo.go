package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"strings"

	"buma-be/internal/domain"
	"buma-be/internal/repository"
)

type debtorRepository struct {
	db *sql.DB
}

func NewDebtorRepository(db *sql.DB) repository.DebtorRepository {
	return &debtorRepository{db: db}
}

func (r *debtorRepository) GetMonitoringData(ctx context.Context, filter domain.DebtorFilter) (*domain.MonitoringResponse, error) {
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 {
		filter.Limit = 10
	}

	// 1. Calculate overall summary metrics
	metricsQuery := `
		SELECT 
			COUNT(d.account_no) as total_debtors,
			COALESCE(SUM(d.outstanding_balance), 0) as total_baki_debet,
			COUNT(CASE WHEN d.exposure_tier = 'Tier 1' THEN 1 END) as tier1_count,
			COUNT(CASE WHEN c.commitment_date = CURRENT_DATE THEN 1 END) as commitments_today
		FROM debtor_accounts d
		LEFT JOIN account_assignments aa ON d.account_no = aa.account_no
		LEFT JOIN officers o ON aa.officer_id = o.officer_id
		LEFT JOIN LATERAL (
			SELECT commitment_date, status, remarks, nominal
			FROM collection_commitments 
			WHERE account_no = d.account_no 
			ORDER BY created_at DESC 
			LIMIT 1
		) c ON true
	`

	var metrics domain.MonitoringMetrics
	err := r.db.QueryRowContext(ctx, metricsQuery).Scan(
		&metrics.TotalDebtors,
		&metrics.TotalBakiDebet,
		&metrics.Tier1Count,
		&metrics.CommitmentsToday,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to fetch monitoring metrics: %w", err)
	}

	if metrics.TotalDebtors > 0 {
		metrics.Tier1Ratio = math.Round((float64(metrics.Tier1Count)/float64(metrics.TotalDebtors)*100)*10) / 10
	}

	// 2. Build filtered and paginated debtor query
	baseQuery := `
		SELECT 
			d.account_no,
			d.debtor_id,
			d.debtor_name,
			d.product_type,
			d.credit_limit,
			d.outstanding_balance,
			d.exposure_tier,
			COALESCE(o.officer_id, '') as officer_id,
			COALESCE(o.full_name, 'Belum Di-assign') as pengelola,
			COALESCE(o.pairing_team, 'Desk Unassigned') as pairing,
			COALESCE(TO_CHAR(c.commitment_date, 'YYYY-MM-DD'), '') as tgl_komitmen,
			COALESCE(c.status, '') as status_komitmen,
			COALESCE(c.remarks, '') as keterangan,
			COALESCE(c.nominal, 0) as nominal_komitmen,
			d.phone,
			d.created_at,
			d.updated_at
		FROM debtor_accounts d
		LEFT JOIN account_assignments aa ON d.account_no = aa.account_no
		LEFT JOIN officers o ON aa.officer_id = o.officer_id
		LEFT JOIN LATERAL (
			SELECT commitment_date, status, remarks, nominal
			FROM collection_commitments 
			WHERE account_no = d.account_no 
			ORDER BY created_at DESC 
			LIMIT 1
		) c ON true
	`

	var whereClauses []string
	var args []interface{}
	argIdx := 1

	if filter.Search != "" {
		searchTerm := "%" + strings.ToLower(filter.Search) + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(LOWER(d.debtor_name) LIKE $%d OR LOWER(d.account_no) LIKE $%d)", argIdx, argIdx))
		args = append(args, searchTerm)
		argIdx++
	}

	if filter.Tier != "" && filter.Tier != "ALL" {
		whereClauses = append(whereClauses, fmt.Sprintf("d.exposure_tier = $%d", argIdx))
		args = append(args, filter.Tier)
		argIdx++
	}

	if filter.OfficerID != "" && filter.OfficerID != "ALL" {
		whereClauses = append(whereClauses, fmt.Sprintf("o.officer_id = $%d", argIdx))
		args = append(args, filter.OfficerID)
		argIdx++
	}

	if filter.OfficerName != "" && filter.OfficerName != "ALL" {
		whereClauses = append(whereClauses, fmt.Sprintf("o.full_name = $%d", argIdx))
		args = append(args, filter.OfficerName)
		argIdx++
	}

	if filter.Product != "" && filter.Product != "ALL" {
		whereClauses = append(whereClauses, fmt.Sprintf("d.product_type = $%d", argIdx))
		args = append(args, filter.Product)
		argIdx++
	}

	whereSQL := ""
	if len(whereClauses) > 0 {
		whereSQL = " WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Count total filtered items
	countQuery := "SELECT COUNT(*) FROM debtor_accounts d LEFT JOIN account_assignments aa ON d.account_no = aa.account_no LEFT JOIN officers o ON aa.officer_id = o.officer_id " + whereSQL
	var totalFiltered int
	err = r.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalFiltered)
	if err != nil {
		return nil, fmt.Errorf("failed to count filtered debtors: %w", err)
	}

	// Sorting
	orderCol := "d.outstanding_balance"
	switch filter.SortBy {
	case "noRekening":
		orderCol = "d.account_no"
	case "namaDebitur":
		orderCol = "d.debtor_name"
	case "limit":
		orderCol = "d.credit_limit"
	case "bakiDebet":
		orderCol = "d.outstanding_balance"
	case "tierEksposur":
		orderCol = "d.exposure_tier"
	case "tglKomitmen":
		orderCol = "c.commitment_date"
	}

	orderDir := "DESC"
	if strings.ToLower(filter.SortDir) == "asc" {
		orderDir = "ASC"
	}

	sortSQL := fmt.Sprintf(" ORDER BY %s %s", orderCol, orderDir)

	// Pagination OFFSET & LIMIT
	offset := (filter.Page - 1) * filter.Limit
	paginationSQL := fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIdx, argIdx+1)
	args = append(args, filter.Limit, offset)

	fullQuery := baseQuery + whereSQL + sortSQL + paginationSQL

	rows, err := r.db.QueryContext(ctx, fullQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute monitoring query: %w", err)
	}
	defer rows.Close()

	var debtors []domain.DebtorAccount
	for rows.Next() {
		var d domain.DebtorAccount
		var tglKomitmen sql.NullString
		var statusKomitmen sql.NullString
		var keterangan sql.NullString
		var nominal sql.NullFloat64

		err := rows.Scan(
			&d.AccountNo,
			&d.DebtorID,
			&d.DebtorName,
			&d.ProductType,
			&d.CreditLimit,
			&d.OutstandingBalance,
			&d.ExposureTier,
			&d.OfficerID,
			&d.OfficerName,
			&d.PairingTeam,
			&tglKomitmen,
			&statusKomitmen,
			&keterangan,
			&nominal,
			&d.Phone,
			&d.CreatedAt,
			&d.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan debtor row: %w", err)
		}

		if tglKomitmen.Valid {
			d.CommitmentDate = tglKomitmen.String
		}
		if statusKomitmen.Valid {
			d.CommitmentStatus = statusKomitmen.String
		}
		if keterangan.Valid {
			d.Remarks = keterangan.String
		}
		if nominal.Valid {
			d.NominalCommitment = nominal.Float64
		}

		debtors = append(debtors, d)
	}

	if debtors == nil {
		debtors = []domain.DebtorAccount{}
	}

	totalPages := 1
	if totalFiltered > 0 {
		totalPages = int(math.Ceil(float64(totalFiltered) / float64(filter.Limit)))
	}

	return &domain.MonitoringResponse{
		Metrics: metrics,
		Data:    debtors,
		Pagination: domain.PaginationMeta{
			CurrentPage:  filter.Page,
			ItemsPerPage: filter.Limit,
			TotalItems:   totalFiltered,
			TotalPages:   totalPages,
		},
	}, nil
}

func (r *debtorRepository) GetDebtorByAccountNo(ctx context.Context, accountNo string) (*domain.DebtorAccount, error) {
	query := `
		SELECT 
			d.account_no,
			d.debtor_id,
			d.debtor_name,
			d.product_type,
			d.credit_limit,
			d.outstanding_balance,
			d.exposure_tier,
			COALESCE(o.officer_id, '') as officer_id,
			COALESCE(o.full_name, 'Belum Di-assign') as pengelola,
			COALESCE(o.pairing_team, 'Desk Unassigned') as pairing,
			COALESCE(TO_CHAR(c.commitment_date, 'YYYY-MM-DD'), '') as tgl_komitmen,
			COALESCE(c.status, '') as status_komitmen,
			COALESCE(c.remarks, '') as keterangan,
			COALESCE(c.nominal, 0) as nominal_komitmen,
			d.phone,
			d.created_at,
			d.updated_at
		FROM debtor_accounts d
		LEFT JOIN account_assignments aa ON d.account_no = aa.account_no
		LEFT JOIN officers o ON aa.officer_id = o.officer_id
		LEFT JOIN LATERAL (
			SELECT commitment_date, status, remarks, nominal
			FROM collection_commitments 
			WHERE account_no = d.account_no 
			ORDER BY created_at DESC 
			LIMIT 1
		) c ON true
		WHERE d.account_no = $1
	`

	var d domain.DebtorAccount
	var tglKomitmen sql.NullString
	var statusKomitmen sql.NullString
	var keterangan sql.NullString
	var nominal sql.NullFloat64

	err := r.db.QueryRowContext(ctx, query, accountNo).Scan(
		&d.AccountNo,
		&d.DebtorID,
		&d.DebtorName,
		&d.ProductType,
		&d.CreditLimit,
		&d.OutstandingBalance,
		&d.ExposureTier,
		&d.OfficerID,
		&d.OfficerName,
		&d.PairingTeam,
		&tglKomitmen,
		&statusKomitmen,
		&keterangan,
		&nominal,
		&d.Phone,
		&d.CreatedAt,
		&d.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get debtor by account_no %s: %w", accountNo, err)
	}

	if tglKomitmen.Valid {
		d.CommitmentDate = tglKomitmen.String
	}
	if statusKomitmen.Valid {
		d.CommitmentStatus = statusKomitmen.String
	}
	if keterangan.Valid {
		d.Remarks = keterangan.String
	}
	if nominal.Valid {
		d.NominalCommitment = nominal.Float64
	}

	return &d, nil
}

func (r *debtorRepository) GetHighExposureTier1Accounts(ctx context.Context) ([]domain.DebtorAccount, error) {
	filter := domain.DebtorFilter{
		Page:  1,
		Limit: 1000,
		Tier:  "Tier 1",
	}
	res, err := r.GetMonitoringData(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}
