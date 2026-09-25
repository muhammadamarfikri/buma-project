package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"buma-be/internal/domain"
	"buma-be/internal/usecase"
	"buma-be/mocks"
)

func TestOfficerUsecase_Mockery(t *testing.T) {
	ctx := context.Background()

	t.Run("CreateOfficer Success", func(t *testing.T) {
		mockOfficerRepo := new(mocks.OfficerRepository)
		uc := usecase.NewOfficerUsecase(mockOfficerRepo)

		req := domain.CreateOfficerRequest{
			OfficerID:   "OFF-005",
			FullName:    "Rizky Ramadhan",
			Email:       "rizky@buma.co.id",
			PairingTeam: "Desk 05",
			Phone:       "081299990000",
		}

		mockOfficerRepo.On("CreateOfficer", ctx, mock.MatchedBy(func(o domain.Officer) bool {
			return o.OfficerID == "OFF-005" && o.FullName == "Rizky Ramadhan" && o.Email == "rizky@buma.co.id"
		})).Return(&domain.Officer{
			OfficerID:   req.OfficerID,
			FullName:    req.FullName,
			Email:       req.Email,
			PairingTeam: req.PairingTeam,
			Phone:       req.Phone,
		}, nil)

		res, err := uc.CreateOfficer(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "OFF-005", res.OfficerID)
		assert.Equal(t, "Rizky Ramadhan", res.FullName)

		mockOfficerRepo.AssertExpectations(t)
	})

	t.Run("CreateOfficer Missing FullName Error", func(t *testing.T) {
		mockOfficerRepo := new(mocks.OfficerRepository)
		uc := usecase.NewOfficerUsecase(mockOfficerRepo)

		req := domain.CreateOfficerRequest{
			Email:       "test@buma.co.id",
			PairingTeam: "Desk 01",
		}

		res, err := uc.CreateOfficer(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "full_name is required", err.Error())
	})

	t.Run("GetAllOfficers Success", func(t *testing.T) {
		mockOfficerRepo := new(mocks.OfficerRepository)
		uc := usecase.NewOfficerUsecase(mockOfficerRepo)

		expectedOfficers := []domain.Officer{
			{OfficerID: "OFF-001", FullName: "Budi Santoso"},
			{OfficerID: "OFF-002", FullName: "Siti Rahma"},
		}

		mockOfficerRepo.On("GetAllOfficers", ctx).Return(expectedOfficers, nil)

		res, err := uc.GetAllOfficers(ctx)

		assert.NoError(t, err)
		assert.Len(t, res, 2)
		assert.Equal(t, "Budi Santoso", res[0].FullName)

		mockOfficerRepo.AssertExpectations(t)
	})

	t.Run("GetOfficerByID Success", func(t *testing.T) {
		mockOfficerRepo := new(mocks.OfficerRepository)
		uc := usecase.NewOfficerUsecase(mockOfficerRepo)

		expectedOfficer := &domain.Officer{
			OfficerID: "OFF-001",
			FullName:  "Budi Santoso",
		}

		mockOfficerRepo.On("GetOfficerByID", ctx, "OFF-001").Return(expectedOfficer, nil)

		res, err := uc.GetOfficerByID(ctx, "OFF-001")

		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "OFF-001", res.OfficerID)

		mockOfficerRepo.AssertExpectations(t)
	})
}
