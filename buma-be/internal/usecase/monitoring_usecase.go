package usecase

import (
	"context"
	"errors"

	"buma-be/internal/domain"
	"buma-be/internal/repository"
)

type MonitoringUsecase interface {
	GetMonitoring(ctx context.Context, filter domain.DebtorFilter) (*domain.MonitoringResponse, error)
	GetDebtorByAccountNo(ctx context.Context, accountNo string) (*domain.DebtorAccount, error)
}

type monitoringUsecase struct {
	debtorRepo repository.DebtorRepository
}

func NewMonitoringUsecase(debtorRepo repository.DebtorRepository) MonitoringUsecase {
	return &monitoringUsecase{debtorRepo: debtorRepo}
}

func (u *monitoringUsecase) GetMonitoring(ctx context.Context, filter domain.DebtorFilter) (*domain.MonitoringResponse, error) {
	return u.debtorRepo.GetMonitoringData(ctx, filter)
}

func (u *monitoringUsecase) GetDebtorByAccountNo(ctx context.Context, accountNo string) (*domain.DebtorAccount, error) {
	if accountNo == "" {
		return nil, errors.New("account_no is required")
	}
	debtor, err := u.debtorRepo.GetDebtorByAccountNo(ctx, accountNo)
	if err != nil {
		return nil, err
	}
	if debtor == nil {
		return nil, errors.New("debtor account not found")
	}
	return debtor, nil
}
