package main

import (
	"log"

	"golang_task_manager_folder_structure/internal/config"
	"golang_task_manager_folder_structure/internal/cron"
	"golang_task_manager_folder_structure/internal/logger"
	"golang_task_manager_folder_structure/internal/repository"
)

func main() {
	// Initialize configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Setup logger
	logger := logger.NewLogger(cfg.LogLevel)
	logger.Info("Starting cron jobs...")

	// Setup database
	db, err := repository.NewDatabase(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	// Initialize repositories
	taskRepo := repository.NewTaskRepository(db)

	// Setup and start scheduler
	scheduler := cron.NewScheduler(taskRepo, logger)
	scheduler.Start()
}
