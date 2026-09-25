package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"

	"buma-be/internal/domain"
	"buma-be/internal/repository"
)

type OfficerUsecase interface {
	CreateOfficer(ctx context.Context, req domain.CreateOfficerRequest) (*domain.Officer, error)
	GetAllOfficers(ctx context.Context) ([]domain.Officer, error)
	GetOfficerByID(ctx context.Context, officerID string) (*domain.Officer, error)
}

type officerUsecase struct {
	officerRepo repository.OfficerRepository
}

func NewOfficerUsecase(officerRepo repository.OfficerRepository) OfficerUsecase {
	return &officerUsecase{officerRepo: officerRepo}
}

func (u *officerUsecase) CreateOfficer(ctx context.Context, req domain.CreateOfficerRequest) (*domain.Officer, error) {
	req.FullName = strings.TrimSpace(req.FullName)
	req.Email = strings.TrimSpace(req.Email)
	req.PairingTeam = strings.TrimSpace(req.PairingTeam)
	req.Phone = strings.TrimSpace(req.Phone)

	if req.FullName == "" {
		return nil, errors.New("full_name is required")
	}
	if req.Email == "" {
		return nil, errors.New("email is required")
	}
	if req.PairingTeam == "" {
		return nil, errors.New("pairing_team is required")
	}

	officerID := strings.TrimSpace(req.OfficerID)
	if officerID == "" {
		officerID = uuid.New().String()
	}

	officer := domain.Officer{
		OfficerID:   officerID,
		FullName:    req.FullName,
		Email:       req.Email,
		PairingTeam: req.PairingTeam,
		Phone:       req.Phone,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return u.officerRepo.CreateOfficer(ctx, officer)
}

func (u *officerUsecase) GetAllOfficers(ctx context.Context) ([]domain.Officer, error) {
	return u.officerRepo.GetAllOfficers(ctx)
}

func (u *officerUsecase) GetOfficerByID(ctx context.Context, officerID string) (*domain.Officer, error) {
	if officerID == "" {
		return nil, errors.New("officer_id parameter is required")
	}
	officer, err := u.officerRepo.GetOfficerByID(ctx, officerID)
	if err != nil {
		return nil, err
	}
	if officer == nil {
		return nil, errors.New("officer not found")
	}
	return officer, nil
}
