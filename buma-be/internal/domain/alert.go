package domain

import "time"

type AlertLog struct {
	ID          string    `json:"id"`
	AccountNo   string    `json:"account_no"`
	AlertType   string    `json:"alert_type"` // UPCOMING_COMMITMENT, OVERDUE_COMMITMENT, HIGH_EXPOSURE_TIER1
	Severity    string    `json:"severity"`   // LOW, MEDIUM, HIGH, CRITICAL
	Message     string    `json:"message"`
	IsProcessed bool      `json:"is_processed"`
	CreatedAt   time.Time `json:"created_at"`
}
