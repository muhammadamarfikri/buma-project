package usecase

import (
	"context"
	"errors"

	"buma-be/internal/domain"
	"buma-be/internal/repository"
	"buma-be/internal/wording"
)

type WordingUsecase interface {
	GenerateWording(ctx context.Context, accountNo string, channel string, tone string) (string, *domain.DebtorAccount, error)
}

type wordingUsecase struct {
	debtorRepo repository.DebtorRepository
}

func NewWordingUsecase(debtorRepo repository.DebtorRepository) WordingUsecase {
	return &wordingUsecase{debtorRepo: debtorRepo}
}

func (u *wordingUsecase) GenerateWording(ctx context.Context, accountNo string, channel string, tone string) (string, *domain.DebtorAccount, error) {
	if accountNo == "" {
		return "", nil, errors.New("account_no query param is required")
	}

	debtor, err := u.debtorRepo.GetDebtorByAccountNo(ctx, accountNo)
	if err != nil {
		return "", nil, err
	}
	if debtor == nil {
		return "", nil, errors.New("debtor account not found")
	}

	text := wording.GenerateCollectionWording(*debtor, channel, tone)
	return text, debtor, nil
}
