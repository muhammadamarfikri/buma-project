package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"strings"

	"buma-be/internal/domain"
	"buma-be/internal/repository"
)

type commitmentRepository struct {
	db *sql.DB
}

func NewCommitmentRepository(db *sql.DB) repository.CommitmentRepository {
	return &commitmentRepository{db: db}
}

func (r *commitmentRepository) CreateCommitment(ctx context.Context, req domain.CreateCommitmentRequest) (*domain.CollectionCommitment, error) {
	nominalVal := req.OutstandingBalance
	if nominalVal <= 0 {
		nominalVal = req.CreditLimit
	}

	query := `
		INSERT INTO collection_commitments (account_no, commitment_date, officer_id, officer_pair_id, status, reason, remarks, nominal, credit_limit, outstanding_balance, exposure_tier, product, debtor_name, created_at, updated_at)
		VALUES ($1, CASE WHEN $2 <> '' THEN $2::DATE ELSE CURRENT_DATE END, $3, $4, $5, $6, $7, $8::NUMERIC, $9::NUMERIC, $10::NUMERIC, $11, $12, $13, NOW(), NOW())
		RETURNING id, account_no, TO_CHAR(commitment_date, 'YYYY-MM-DD'), COALESCE(officer_id, ''), COALESCE(officer_pair_id, ''), status, COALESCE(reason, ''), COALESCE(remarks, ''), COALESCE(nominal, 0::NUMERIC), COALESCE(credit_limit, 0::NUMERIC), COALESCE(outstanding_balance, 0::NUMERIC), COALESCE(exposure_tier, ''), COALESCE(product, ''), COALESCE(debtor_name, ''), created_at, updated_at
	`

	var c domain.CollectionCommitment
	err := r.db.QueryRowContext(
		ctx,
		query,
		req.AccountNo,
		req.CommitmentDate,
		req.OfficerID,
		req.OfficerPairID,
		req.Status,
		req.Reason,
		req.Remarks,
		nominalVal,
		req.CreditLimit,
		req.OutstandingBalance,
		req.ExposureTier,
		req.Product,
		req.DebtorName,
	).Scan(
		&c.ID,
		&c.AccountNo,
		&c.CommitmentDate,
		&c.OfficerID,
		&c.OfficerPairID,
		&c.Status,
		&c.Reason,
		&c.Remarks,
		&c.Nominal,
		&c.CreditLimit,
		&c.OutstandingBalance,
		&c.ExposureTier,
		&c.Product,
		&c.DebtorName,
		&c.CreatedAt,
		&c.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert commitment record: %w", err)
	}

	upsertDebtorQuery := `
		INSERT INTO debtor_accounts (account_no, debtor_id, debtor_name, product_type, credit_limit, outstanding_balance, exposure_tier, phone, created_at, updated_at)
		VALUES ($1, 'D-' || $1, COALESCE(NULLIF($2, ''), 'Debtor ' || $1), COALESCE(NULLIF($3, ''), 'KMK'), $4::NUMERIC, $5::NUMERIC, COALESCE(NULLIF($6, ''), 'Tier 1'), '-', NOW(), NOW())
		ON CONFLICT (account_no) DO UPDATE 
		SET debtor_name = COALESCE(NULLIF(EXCLUDED.debtor_name, ''), debtor_accounts.debtor_name),
		    credit_limit = CASE WHEN EXCLUDED.credit_limit > 0 THEN EXCLUDED.credit_limit ELSE debtor_accounts.credit_limit END,
		    outstanding_balance = CASE WHEN EXCLUDED.outstanding_balance > 0 THEN EXCLUDED.outstanding_balance ELSE debtor_accounts.outstanding_balance END,
		    exposure_tier = COALESCE(NULLIF(EXCLUDED.exposure_tier, ''), debtor_accounts.exposure_tier),
		    updated_at = NOW()
	`
	_, _ = r.db.ExecContext(ctx, upsertDebtorQuery, req.AccountNo, req.DebtorName, req.Product, req.CreditLimit, req.OutstandingBalance, req.ExposureTier)

	return &c, nil
}

func (r *commitmentRepository) UpdateCommitmentByReason(ctx context.Context, req domain.CreateCommitmentRequest) (*domain.CollectionCommitment, error) {
	nominalVal := req.OutstandingBalance
	if nominalVal <= 0 {
		nominalVal = req.CreditLimit
	}

	query := `
		UPDATE collection_commitments
		SET commitment_date = CASE WHEN $1 <> '' THEN $1::DATE ELSE commitment_date END,
		    officer_id = COALESCE(NULLIF($2, ''), officer_id),
		    officer_pair_id = COALESCE(NULLIF($3, ''), officer_pair_id),
		    status = COALESCE(NULLIF($4, ''), status),
		    remarks = COALESCE(NULLIF($5, ''), remarks),
		    nominal = CASE WHEN $6::NUMERIC > 0::NUMERIC THEN $6::NUMERIC ELSE nominal END,
		    credit_limit = CASE WHEN $7::NUMERIC > 0::NUMERIC THEN $7::NUMERIC ELSE credit_limit END,
		    outstanding_balance = CASE WHEN $8::NUMERIC > 0::NUMERIC THEN $8::NUMERIC ELSE outstanding_balance END,
		    exposure_tier = COALESCE(NULLIF($9, ''), exposure_tier),
		    product = COALESCE(NULLIF($10, ''), product),
		    debtor_name = COALESCE(NULLIF($11, ''), debtor_name),
		    updated_at = NOW()
		WHERE account_no = $12 AND LOWER(TRIM(reason)) = LOWER(TRIM($13))
		RETURNING id, account_no, TO_CHAR(commitment_date, 'YYYY-MM-DD'), COALESCE(officer_id, ''), COALESCE(officer_pair_id, ''), status, COALESCE(reason, ''), COALESCE(remarks, ''), COALESCE(nominal, 0::NUMERIC), COALESCE(credit_limit, 0::NUMERIC), COALESCE(outstanding_balance, 0::NUMERIC), COALESCE(exposure_tier, ''), COALESCE(product, ''), COALESCE(debtor_name, ''), created_at, updated_at
	`

	var c domain.CollectionCommitment
	err := r.db.QueryRowContext(
		ctx,
		query,
		req.CommitmentDate,
		req.OfficerID,
		req.OfficerPairID,
		req.Status,
		req.Remarks,
		nominalVal,
		req.CreditLimit,
		req.OutstandingBalance,
		req.ExposureTier,
		req.Product,
		req.DebtorName,
		req.AccountNo,
		req.Reason,
	).Scan(
		&c.ID,
		&c.AccountNo,
		&c.CommitmentDate,
		&c.OfficerID,
		&c.OfficerPairID,
		&c.Status,
		&c.Reason,
		&c.Remarks,
		&c.Nominal,
		&c.CreditLimit,
		&c.OutstandingBalance,
		&c.ExposureTier,
		&c.Product,
		&c.DebtorName,
		&c.CreatedAt,
		&c.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update commitment record for reason '%s': %w", req.Reason, err)
	}

	upsertDebtorQuery := `
		INSERT INTO debtor_accounts (account_no, debtor_id, debtor_name, product_type, credit_limit, outstanding_balance, exposure_tier, phone, created_at, updated_at)
		VALUES ($1, 'D-' || $1, COALESCE(NULLIF($2, ''), 'Debtor ' || $1), COALESCE(NULLIF($3, ''), 'KMK'), $4::NUMERIC, $5::NUMERIC, COALESCE(NULLIF($6, ''), 'Tier 1'), '-', NOW(), NOW())
		ON CONFLICT (account_no) DO UPDATE 
		SET debtor_name = COALESCE(NULLIF(EXCLUDED.debtor_name, ''), debtor_accounts.debtor_name),
		    credit_limit = CASE WHEN EXCLUDED.credit_limit > 0 THEN EXCLUDED.credit_limit ELSE debtor_accounts.credit_limit END,
		    outstanding_balance = CASE WHEN EXCLUDED.outstanding_balance > 0 THEN EXCLUDED.outstanding_balance ELSE debtor_accounts.outstanding_balance END,
		    exposure_tier = COALESCE(NULLIF(EXCLUDED.exposure_tier, ''), debtor_accounts.exposure_tier),
		    updated_at = NOW()
	`
	_, _ = r.db.ExecContext(ctx, upsertDebtorQuery, req.AccountNo, req.DebtorName, req.Product, req.CreditLimit, req.OutstandingBalance, req.ExposureTier)

	return &c, nil
}

func (r *commitmentRepository) CreateOrUpdateCommitment(ctx context.Context, req domain.CreateCommitmentRequest) (*domain.CollectionCommitment, error) {
	updated, err := r.UpdateCommitmentByReason(ctx, req)
	if err == nil && updated != nil {
		return updated, nil
	}
	return r.CreateCommitment(ctx, req)
}

func (r *commitmentRepository) GetCommitmentsByAccountNo(ctx context.Context, accountNo string) ([]domain.CollectionCommitment, error) {
	query := `
		SELECT id, account_no, TO_CHAR(commitment_date, 'YYYY-MM-DD'), COALESCE(officer_id, ''), COALESCE(officer_pair_id, ''), status, COALESCE(reason, ''), COALESCE(remarks, ''), COALESCE(nominal, 0), created_at, updated_at
		FROM collection_commitments
		WHERE account_no = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, accountNo)
	if err != nil {
		return nil, fmt.Errorf("failed to query commitments for account_no %s: %w", accountNo, err)
	}
	defer rows.Close()

	var list []domain.CollectionCommitment
	for rows.Next() {
		var c domain.CollectionCommitment
		if err := rows.Scan(&c.ID, &c.AccountNo, &c.CommitmentDate, &c.OfficerID, &c.OfficerPairID, &c.Status, &c.Reason, &c.Remarks, &c.Nominal, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, c)
	}

	if list == nil {
		list = []domain.CollectionCommitment{}
	}

	return list, nil
}

func (r *commitmentRepository) GetCommitmentsByDate(ctx context.Context, dateStr string) ([]domain.CollectionCommitment, error) {
	query := `
		SELECT id, account_no, TO_CHAR(commitment_date, 'YYYY-MM-DD'), COALESCE(officer_id, ''), COALESCE(officer_pair_id, ''), status, COALESCE(reason, ''), COALESCE(remarks, ''), COALESCE(nominal, 0), created_at, updated_at
		FROM collection_commitments
		WHERE commitment_date = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, dateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to query commitments for date %s: %w", dateStr, err)
	}
	defer rows.Close()

	var list []domain.CollectionCommitment
	for rows.Next() {
		var c domain.CollectionCommitment
		if err := rows.Scan(&c.ID, &c.AccountNo, &c.CommitmentDate, &c.OfficerID, &c.OfficerPairID, &c.Status, &c.Reason, &c.Remarks, &c.Nominal, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		list = append(list, c)
	}

	if list == nil {
		list = []domain.CollectionCommitment{}
	}

	return list, nil
}

func (r *commitmentRepository) GetCommitments(ctx context.Context, filter domain.CommitmentFilter) (*domain.PaginatedCommitmentResponse, error) {
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.Limit < 1 {
		filter.Limit = 10
	}

	whereClauses := []string{"1=1"}
	args := []interface{}{}
	argIdx := 1

	if filter.AccountNo != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.account_no = $%d", argIdx))
		args = append(args, filter.AccountNo)
		argIdx++
	}

	if filter.Status != "" && filter.Status != "ALL" {
		whereClauses = append(whereClauses, fmt.Sprintf("c.status = $%d", argIdx))
		args = append(args, filter.Status)
		argIdx++
	}

	if filter.Tier != "" && filter.Tier != "ALL" {
		whereClauses = append(whereClauses, fmt.Sprintf("(c.exposure_tier = $%d OR d.exposure_tier = $%d)", argIdx, argIdx))
		args = append(args, filter.Tier)
		argIdx++
	}

	if filter.Product != "" && filter.Product != "ALL" {
		whereClauses = append(whereClauses, fmt.Sprintf("(c.product = $%d OR d.product_type = $%d)", argIdx, argIdx))
		args = append(args, filter.Product)
		argIdx++
	}

	if filter.Pengelola != "" && filter.Pengelola != "ALL" {
		whereClauses = append(whereClauses, fmt.Sprintf("(c.officer_id = $%d OR o.full_name ILIKE $%d)", argIdx, argIdx+1))
		args = append(args, filter.Pengelola, "%"+filter.Pengelola+"%")
		argIdx += 2
	}

	if filter.Search != "" {
		searchTerm := "%" + filter.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(c.account_no ILIKE $%d OR c.status ILIKE $%d OR COALESCE(c.reason,'') ILIKE $%d OR COALESCE(c.remarks,'') ILIKE $%d OR COALESCE(d.debtor_name,'') ILIKE $%d OR COALESCE(o.full_name,'') ILIKE $%d)", argIdx, argIdx, argIdx, argIdx, argIdx, argIdx))
		args = append(args, searchTerm)
		argIdx++
	}

	whereStmt := strings.Join(whereClauses, " AND ")

	var totalItems int
	if r.db != nil {
		countQuery := fmt.Sprintf(`
			SELECT COUNT(*) 
			FROM collection_commitments c 
			LEFT JOIN debtor_accounts d ON c.account_no = d.account_no 
			LEFT JOIN officers o ON c.officer_id = o.officer_id 
			WHERE %s
		`, whereStmt)
		err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalItems)
		if err != nil {
			return nil, fmt.Errorf("failed to count commitments: %w", err)
		}
	}

	totalPages := 1
	if totalItems > 0 {
		totalPages = int(math.Ceil(float64(totalItems) / float64(filter.Limit)))
	}

	offset := (filter.Page - 1) * filter.Limit

	orderCol := "c.created_at"
	switch filter.SortBy {
	case "account_no", "noRekening":
		orderCol = "c.account_no"
	case "commitment_date", "tglKomitmen":
		orderCol = "c.commitment_date"
	case "status", "statusKomitmen":
		orderCol = "c.status"
	case "tier", "exposure_tier", "tierEksposur":
		orderCol = "COALESCE(c.exposure_tier, d.exposure_tier)"
	case "product", "produk":
		orderCol = "COALESCE(c.product, d.product_type)"
	}

	orderDir := "DESC"
	if strings.ToLower(filter.SortDir) == "asc" {
		orderDir = "ASC"
	}
	orderStmt := fmt.Sprintf(" ORDER BY %s %s", orderCol, orderDir)

	var list []domain.CollectionCommitment
	if r.db != nil {
		dataQuery := fmt.Sprintf(`
			SELECT c.id, c.account_no, TO_CHAR(c.commitment_date, 'YYYY-MM-DD'), COALESCE(c.officer_id, ''), COALESCE(c.officer_pair_id, ''), c.status, COALESCE(c.reason, ''), COALESCE(c.remarks, ''), COALESCE(c.nominal, 0), COALESCE(c.credit_limit, 0), COALESCE(c.outstanding_balance, 0), COALESCE(c.exposure_tier, ''), COALESCE(c.product, ''), c.created_at, c.updated_at
			FROM collection_commitments c
			LEFT JOIN debtor_accounts d ON c.account_no = d.account_no
			LEFT JOIN officers o ON c.officer_id = o.officer_id
			WHERE %s
			%s
			LIMIT $%d OFFSET $%d
		`, whereStmt, orderStmt, argIdx, argIdx+1)

		queryArgs := append(args, filter.Limit, offset)
		rows, err := r.db.QueryContext(ctx, dataQuery, queryArgs...)
		if err != nil {
			return nil, fmt.Errorf("failed to query paginated commitments: %w", err)
		}
		defer rows.Close()

		for rows.Next() {
			var c domain.CollectionCommitment
			if err := rows.Scan(&c.ID, &c.AccountNo, &c.CommitmentDate, &c.OfficerID, &c.OfficerPairID, &c.Status, &c.Reason, &c.Remarks, &c.Nominal, &c.CreditLimit, &c.OutstandingBalance, &c.ExposureTier, &c.Product, &c.CreatedAt, &c.UpdatedAt); err != nil {
				return nil, err
			}
			list = append(list, c)
		}
	}

	if list == nil {
		list = []domain.CollectionCommitment{}
	}

	return &domain.PaginatedCommitmentResponse{
		Data: list,
		Pagination: domain.PaginationMeta{
			CurrentPage:  filter.Page,
			ItemsPerPage: filter.Limit,
			TotalItems:   totalItems,
			TotalPages:   totalPages,
		},
	}, nil
}

func (r *commitmentRepository) GetCommitmentByID(ctx context.Context, id string) (*domain.CollectionCommitment, error) {
	if r.db == nil {
		return nil, fmt.Errorf("database connection is not available")
	}

	query := `
		SELECT id, account_no, TO_CHAR(commitment_date, 'YYYY-MM-DD'), COALESCE(officer_id, ''), COALESCE(officer_pair_id, ''), status, COALESCE(reason, ''), COALESCE(remarks, ''), COALESCE(nominal, 0), created_at, updated_at
		FROM collection_commitments
		WHERE id = $1
	`

	var c domain.CollectionCommitment
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&c.ID, &c.AccountNo, &c.CommitmentDate, &c.OfficerID, &c.OfficerPairID, &c.Status, &c.Reason, &c.Remarks, &c.Nominal, &c.CreatedAt, &c.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("commitment with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to get commitment by id %s: %w", id, err)
	}

	return &c, nil
}

func (r *commitmentRepository) UpdateCommitmentByID(ctx context.Context, id string, req domain.UpdateCommitmentRequest) (*domain.CollectionCommitment, error) {
	if r.db == nil {
		return nil, fmt.Errorf("database connection is not available")
	}

	nominalVal := req.Nominal

	query := `
		UPDATE collection_commitments
		SET commitment_date = CASE WHEN $1 <> '' THEN $1::DATE ELSE commitment_date END,
			status = COALESCE(NULLIF($2, ''), status),
			reason = COALESCE(NULLIF($3, ''), reason),
			remarks = COALESCE(NULLIF($4, ''), remarks),
			nominal = $5::NUMERIC,
			officer_id = COALESCE(NULLIF($6, ''), officer_id),
			officer_pair_id = COALESCE(NULLIF($7, ''), officer_pair_id),
			credit_limit = CASE WHEN $8::NUMERIC > 0::NUMERIC THEN $8::NUMERIC ELSE credit_limit END,
			outstanding_balance = CASE WHEN $9::NUMERIC > 0::NUMERIC THEN $9::NUMERIC ELSE outstanding_balance END,
			exposure_tier = COALESCE(NULLIF($10, ''), exposure_tier),
			product = COALESCE(NULLIF($11, ''), product),
			updated_at = NOW()
		WHERE id::text = $12 OR account_no = $12
		RETURNING id, account_no, TO_CHAR(commitment_date, 'YYYY-MM-DD'), COALESCE(officer_id, ''), COALESCE(officer_pair_id, ''), status, COALESCE(reason, ''), COALESCE(remarks, ''), COALESCE(nominal, 0::NUMERIC), COALESCE(credit_limit, 0::NUMERIC), COALESCE(outstanding_balance, 0::NUMERIC), COALESCE(exposure_tier, ''), COALESCE(product, ''), created_at, updated_at
	`

	var c domain.CollectionCommitment
	err := r.db.QueryRowContext(
		ctx,
		query,
		req.CommitmentDate,
		req.Status,
		req.Reason,
		req.Remarks,
		nominalVal,
		req.OfficerID,
		req.OfficerPairID,
		req.CreditLimit,
		req.OutstandingBalance,
		req.ExposureTier,
		req.Product,
		id,
	).Scan(
		&c.ID, &c.AccountNo, &c.CommitmentDate, &c.OfficerID, &c.OfficerPairID, &c.Status, &c.Reason, &c.Remarks, &c.Nominal, &c.CreditLimit, &c.OutstandingBalance, &c.ExposureTier, &c.Product, &c.CreatedAt, &c.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to update commitment with id %s: %w", id, err)
	}

	updateDebtorQuery := `
		UPDATE debtor_accounts 
		SET credit_limit = CASE WHEN $1::NUMERIC > 0::NUMERIC THEN $1::NUMERIC ELSE credit_limit END,
		    outstanding_balance = CASE WHEN $2::NUMERIC > 0::NUMERIC THEN $2::NUMERIC ELSE outstanding_balance END,
		    exposure_tier = COALESCE(NULLIF($3, ''), exposure_tier),
		    updated_at = NOW()
		WHERE account_no = $4
	`
	_, _ = r.db.ExecContext(ctx, updateDebtorQuery, req.CreditLimit, req.OutstandingBalance, req.ExposureTier, c.AccountNo)

	if req.OfficerID != "" {
		_, _ = r.db.ExecContext(ctx, `
			INSERT INTO account_assignments (account_no, officer_id, assigned_at)
			VALUES ($1, $2, NOW())
			ON CONFLICT (account_no) DO UPDATE SET officer_id = $2, assigned_at = NOW()
		`, c.AccountNo, req.OfficerID)
	}

	return &c, nil
}
