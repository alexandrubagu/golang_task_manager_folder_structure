package cron

import (
	"time"

	"golang_task_manager_folder_structure/internal/logger"
	"golang_task_manager_folder_structure/internal/repository"

	"github.com/go-co-op/gocron"
)

// Scheduler runs recurring jobs
type Scheduler struct {
	scheduler *gocron.Scheduler
	repo      *repository.TaskRepository
	logger    *logger.Logger
}

// NewScheduler creates a new scheduler
func NewScheduler(repo *repository.TaskRepository, logger *logger.Logger) *Scheduler {
	s := gocron.NewScheduler(time.UTC)

	return &Scheduler{
		scheduler: s,
		repo:      repo,
		logger:    logger,
	}
}

// Start begins the scheduler
func (s *Scheduler) Start() {
	// Schedule task reminder job to run daily at 9 AM
	s.scheduler.Every(1).Day().At("09:00").Do(func() {
		TaskReminder(s.repo, s.logger)
	})

	// Schedule cleanup job to run weekly on Sunday at midnight
	s.scheduler.Every(1).Week().Sunday().At("00:00").Do(func() {
		CleanupOldTasks(s.repo, s.logger)
	})

	// Start scheduler
	s.scheduler.StartBlocking()
}
