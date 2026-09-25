package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"buma-be/internal/domain"
	"buma-be/internal/repository"
)

type officerRepository struct {
	db *sql.DB
}

func NewOfficerRepository(db *sql.DB) repository.OfficerRepository {
	return &officerRepository{db: db}
}

func (r *officerRepository) CreateOfficer(ctx context.Context, officer domain.Officer) (*domain.Officer, error) {
	query := `
		INSERT INTO officers (officer_id, full_name, email, pairing_team, phone, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING officer_id, full_name, email, pairing_team, COALESCE(phone, ''), created_at, updated_at
	`

	var o domain.Officer
	err := r.db.QueryRowContext(
		ctx,
		query,
		officer.OfficerID,
		officer.FullName,
		officer.Email,
		officer.PairingTeam,
		officer.Phone,
	).Scan(
		&o.OfficerID,
		&o.FullName,
		&o.Email,
		&o.PairingTeam,
		&o.Phone,
		&o.CreatedAt,
		&o.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert officer: %w", err)
	}

	return &o, nil
}

func (r *officerRepository) GetAllOfficers(ctx context.Context) ([]domain.Officer, error) {
	query := `
		SELECT officer_id, full_name, email, pairing_team, COALESCE(phone, ''), created_at, updated_at
		FROM officers
		ORDER BY full_name ASC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query officers: %w", err)
	}
	defer rows.Close()

	var officers []domain.Officer
	for rows.Next() {
		var o domain.Officer
		if err := rows.Scan(&o.OfficerID, &o.FullName, &o.Email, &o.PairingTeam, &o.Phone, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, err
		}
		officers = append(officers, o)
	}

	if officers == nil {
		officers = []domain.Officer{}
	}

	return officers, nil
}

func (r *officerRepository) GetOfficerByID(ctx context.Context, officerID string) (*domain.Officer, error) {
	query := `
		SELECT officer_id, full_name, email, pairing_team, COALESCE(phone, ''), created_at, updated_at
		FROM officers
		WHERE officer_id = $1
	`

	var o domain.Officer
	err := r.db.QueryRowContext(ctx, query, officerID).Scan(&o.OfficerID, &o.FullName, &o.Email, &o.PairingTeam, &o.Phone, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query officer by id %s: %w", officerID, err)
	}

	return &o, nil
}
