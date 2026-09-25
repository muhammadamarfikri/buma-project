package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"buma-be/internal/domain"
	"buma-be/internal/repository"
)

type alertRepository struct {
	db *sql.DB
}

func NewAlertRepository(db *sql.DB) repository.AlertRepository {
	return &alertRepository{db: db}
}

func (r *alertRepository) CreateAlertLog(ctx context.Context, alert domain.AlertLog) error {
	query := `
		INSERT INTO alert_logs (account_no, alert_type, severity, message, is_processed, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
	`
	_, err := r.db.ExecContext(ctx, query, alert.AccountNo, alert.AlertType, alert.Severity, alert.Message, alert.IsProcessed)
	if err != nil {
		return fmt.Errorf("failed to create alert log: %w", err)
	}
	return nil
}

func (r *alertRepository) GetUnprocessedAlerts(ctx context.Context) ([]domain.AlertLog, error) {
	query := `
		SELECT id, account_no, alert_type, severity, message, is_processed, created_at
		FROM alert_logs
		WHERE is_processed = false
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch unprocessed alert logs: %w", err)
	}
	defer rows.Close()

	var alerts []domain.AlertLog
	for rows.Next() {
		var a domain.AlertLog
		if err := rows.Scan(&a.ID, &a.AccountNo, &a.AlertType, &a.Severity, &a.Message, &a.IsProcessed, &a.CreatedAt); err != nil {
			return nil, err
		}
		alerts = append(alerts, a)
	}

	if alerts == nil {
		alerts = []domain.AlertLog{}
	}

	return alerts, nil
}
