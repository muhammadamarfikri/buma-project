package scheduler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"

	"buma-be/internal/domain"
	"buma-be/internal/repository"
)

type AlertScheduler struct {
	cron       *cron.Cron
	debtorRepo repository.DebtorRepository
	alertRepo  repository.AlertRepository
}

func NewAlertScheduler(debtorRepo repository.DebtorRepository, alertRepo repository.AlertRepository) *AlertScheduler {
	c := cron.New(cron.WithSeconds())
	return &AlertScheduler{
		cron:       c,
		debtorRepo: debtorRepo,
		alertRepo:  alertRepo,
	}
}

func (s *AlertScheduler) Start() {
	// Schedule job to run every hour (e.g. "0 0 * * * *" or cron spec "@hourly")
	_, err := s.cron.AddFunc("@hourly", func() {
		s.RunAlertCheck(context.Background())
	})
	if err != nil {
		log.Printf("[Scheduler] Error scheduling alert check job: %v\n", err)
		return
	}

	s.cron.Start()
	log.Println("[Scheduler] Alert & Debt Collection Scheduler Engine started successfully")

	// Trigger initial check on startup
	go s.RunAlertCheck(context.Background())
}

func (s *AlertScheduler) Stop() {
	s.cron.Stop()
	log.Println("[Scheduler] Alert Scheduler stopped")
}

func (s *AlertScheduler) RunAlertCheck(ctx context.Context) {
	log.Println("[Scheduler] Running automated debt collection risk check...")

	// 1. Scan Tier 1 high-exposure accounts
	tier1Accounts, err := s.debtorRepo.GetHighExposureTier1Accounts(ctx)
	if err != nil {
		log.Printf("[Scheduler] Error querying Tier 1 accounts: %v\n", err)
	} else {
		for _, account := range tier1Accounts {
			msg := fmt.Sprintf("High Exposure Facility Alert: Account %s (%s) has high exposure %s with outstanding balance Rp %.0f. Immediate officer review required.",
				account.AccountNo, account.DebtorName, account.ExposureTier, account.OutstandingBalance)

			alert := domain.AlertLog{
				AccountNo:   account.AccountNo,
				AlertType:   "HIGH_EXPOSURE_TIER1",
				Severity:    "HIGH",
				Message:     msg,
				IsProcessed: false,
			}
			if err := s.alertRepo.CreateAlertLog(ctx, alert); err != nil {
				log.Printf("[Scheduler] Error creating alert log for account %s: %v\n", account.AccountNo, err)
			}
		}
		log.Printf("[Scheduler] Risk evaluation completed. Checked %d Tier 1 high-exposure debtor accounts.\n", len(tier1Accounts))
	}

	// 2. Scan today's commitments
	todayStr := time.Now().Format("2006-01-02")
	filter := domain.DebtorFilter{Page: 1, Limit: 1000}
	res, err := s.debtorRepo.GetMonitoringData(ctx, filter)
	if err == nil && res != nil {
		for _, account := range res.Data {
			if account.CommitmentDate == todayStr {
				msg := fmt.Sprintf("Commitment Due Today Alert: Account %s (%s) commitment date is TODAY (%s). Status: %s. Officer %s assigned.",
					account.AccountNo, account.DebtorName, todayStr, account.CommitmentStatus, account.OfficerName)

				alert := domain.AlertLog{
					AccountNo:   account.AccountNo,
					AlertType:   "UPCOMING_COMMITMENT",
					Severity:    "MEDIUM",
					Message:     msg,
					IsProcessed: false,
				}
				_ = s.alertRepo.CreateAlertLog(ctx, alert)
			}
		}
	}
}
