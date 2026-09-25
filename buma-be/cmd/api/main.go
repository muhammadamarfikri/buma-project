package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"buma-be/config"
	"buma-be/internal/handler"
	"buma-be/internal/repository"
	"buma-be/internal/repository/memory"
	"buma-be/internal/repository/postgres"
	"buma-be/internal/router"
	"buma-be/internal/scheduler"
	"buma-be/internal/usecase"
)

func main() {
	log.Println("==========================================================================")
	log.Println("Starting BUMA Automated Debt Collection & Reminder Dispatch System Backend")
	log.Println("==========================================================================")

	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Initialize Database Connection (or Fallback Memory Mode)
	var debtorRepo repository.DebtorRepository
	var commitmentRepo repository.CommitmentRepository
	var officerRepo repository.OfficerRepository
	var alertRepo repository.AlertRepository

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	var dbConnected bool
	if err == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		if err := db.PingContext(ctx); err == nil {
			dbConnected = true
			log.Println("[DB] Successfully connected to PostgreSQL database.")
		}
		cancel()
	}

	if dbConnected {
		debtorRepo = postgres.NewDebtorRepository(db)
		commitmentRepo = postgres.NewCommitmentRepository(db)
		officerRepo = postgres.NewOfficerRepository(db)
		alertRepo = postgres.NewAlertRepository(db)
	} else {
		log.Println("[DB Warning] Could not reach active PostgreSQL database instance.")
		log.Println("[DB Info] Operating in High-Availability In-Memory mode with full BUMA debtor seed dataset.")
		debtorRepo = memory.NewInMemoryDebtorRepository()
		commitmentRepo = memory.NewInMemoryCommitmentRepository()
		officerRepo = memory.NewInMemoryOfficerRepository()
		alertRepo = postgres.NewAlertRepository(db)
	}

	// 3. Initialize Usecases / Business Logic Services
	monitoringUc := usecase.NewMonitoringUsecase(debtorRepo)
	commitmentUc := usecase.NewCommitmentUsecase(commitmentRepo, debtorRepo, officerRepo)
	wordingUc := usecase.NewWordingUsecase(debtorRepo)
	officerUc := usecase.NewOfficerUsecase(officerRepo)

	// 4. Initialize HTTP Handlers
	monitoringHnd := handler.NewMonitoringHandler(monitoringUc)
	commitmentHnd := handler.NewCommitmentHandler(commitmentUc)
	wordingHnd := handler.NewWordingHandler(wordingUc)
	officerHnd := handler.NewOfficerHandler(officerUc)

	// 5. Initialize & Start Background Alert Scheduler
	alertSched := scheduler.NewAlertScheduler(debtorRepo, alertRepo)
	alertSched.Start()
	defer alertSched.Stop()

	// 6. Setup Echo Router
	routerDeps := router.RouterDependencies{
		MonitoringHandler: monitoringHnd,
		CommitmentHandler: commitmentHnd,
		WordingHandler:    wordingHnd,
		OfficerHandler:    officerHnd,
	}
	e := router.SetupRouter(routerDeps)

	// 7. Start Echo HTTP Server with Graceful Shutdown
	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      e,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("[Server] BUMA Backend API server (Echo) listening on http://localhost:%s\n", cfg.Port)
		log.Printf("[Server] Health check: http://localhost:%s/health\n", cfg.Port)
		log.Printf("[Server] Monitoring API: http://localhost:%s/api/v1/monitoring\n", cfg.Port)
		log.Printf("[Server] Wording API: http://localhost:%s/api/v1/templates/wording?account_no=1029384756\n", cfg.Port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[Server Fatal] HTTP server failed to start: %v\n", err)
		}
	}()

	// 8. Graceful Shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("[Server] Shutting down BUMA Backend API server gracefully...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("[Server Warning] Forced server shutdown: %v\n", err)
	}

	if db != nil {
		_ = db.Close()
	}

	log.Println("[Server] BUMA Backend API server stopped clean.")
}
