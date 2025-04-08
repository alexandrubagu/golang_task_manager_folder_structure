package cron

import (
	"time"

	"golang_task_manager_folder_structure/internal/logger"
	"golang_task_manager_folder_structure/internal/repository"
)

// TaskReminder checks for upcoming tasks and potentially sends reminders
func TaskReminder(repo *repository.TaskRepository, log *logger.Logger) {
	log.Info("Running task reminder job")

	// Get all incomplete tasks
	tasks, err := repo.FindAll()
	if err != nil {
		log.Error("Failed to get tasks for reminder", err)
		return
	}

	tomorrow := time.Now().AddDate(0, 0, 1)

	for _, task := range tasks {
		if !task.Completed && task.DueDate != nil && task.DueDate.Before(tomorrow) {
			log.Info("Task #%d '%s' is due soon (%s)", task.ID, task.Title, task.DueDate.Format("2006-01-02"))
			// In a real application, we would send an email or notification here
		}
	}
}

// CleanupOldTasks archives or removes old completed tasks
func CleanupOldTasks(repo *repository.TaskRepository, log *logger.Logger) {
	log.Info("Running cleanup job for old tasks")

	// In a real application, we would archive or delete old tasks
	// This is just a placeholder
	log.Info("Cleanup job completed")
}
