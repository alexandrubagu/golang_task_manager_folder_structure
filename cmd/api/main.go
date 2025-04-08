package main

import (
	"log"

	"golang_task_manager_folder_structure/internal/api"
	"golang_task_manager_folder_structure/internal/config"
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
	logger.Info("Starting API server...")

	// Setup database
	db, err := repository.NewDatabase(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	// Initialize repositories
	taskRepo := repository.NewTaskRepository(db)

	// Initialize services
	services := api.NewServices(taskRepo, logger)

	// Setup and start server
	server := api.NewServer(cfg, services, logger)
	if err := server.Start(); err != nil {
		logger.Fatal("Failed to start server", err)
	}
}
