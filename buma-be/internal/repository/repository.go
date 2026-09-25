package repository

import (
	"context"

	"buma-be/internal/domain"
)

type DebtorRepository interface {
	GetMonitoringData(ctx context.Context, filter domain.DebtorFilter) (*domain.MonitoringResponse, error)
	GetDebtorByAccountNo(ctx context.Context, accountNo string) (*domain.DebtorAccount, error)
	GetHighExposureTier1Accounts(ctx context.Context) ([]domain.DebtorAccount, error)
}

type CommitmentRepository interface {
	CreateCommitment(ctx context.Context, req domain.CreateCommitmentRequest) (*domain.CollectionCommitment, error)
	UpdateCommitmentByReason(ctx context.Context, req domain.CreateCommitmentRequest) (*domain.CollectionCommitment, error)
	CreateOrUpdateCommitment(ctx context.Context, req domain.CreateCommitmentRequest) (*domain.CollectionCommitment, error)
	GetCommitmentsByAccountNo(ctx context.Context, accountNo string) ([]domain.CollectionCommitment, error)
	GetCommitmentsByDate(ctx context.Context, date string) ([]domain.CollectionCommitment, error)
	GetCommitments(ctx context.Context, filter domain.CommitmentFilter) (*domain.PaginatedCommitmentResponse, error)
	GetCommitmentByID(ctx context.Context, id string) (*domain.CollectionCommitment, error)
	UpdateCommitmentByID(ctx context.Context, id string, req domain.UpdateCommitmentRequest) (*domain.CollectionCommitment, error)
}

type OfficerRepository interface {
	CreateOfficer(ctx context.Context, officer domain.Officer) (*domain.Officer, error)
	GetAllOfficers(ctx context.Context) ([]domain.Officer, error)
	GetOfficerByID(ctx context.Context, officerID string) (*domain.Officer, error)
}

type AlertRepository interface {
	CreateAlertLog(ctx context.Context, alert domain.AlertLog) error
	GetUnprocessedAlerts(ctx context.Context) ([]domain.AlertLog, error)
}
