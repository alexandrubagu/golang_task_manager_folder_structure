package repository

import (
	"database/sql"
	"errors"
	"time"
)

// Task represents a task entity
type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// TaskRepository handles DB operations for tasks
type TaskRepository struct {
	db *sql.DB
}

// NewTaskRepository creates a new TaskRepository
func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

// Initialize creates task table if it doesn't exist
func (r *TaskRepository) Initialize() error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		completed BOOLEAN DEFAULT FALSE,
		due_date DATETIME,
		completed_at DATETIME,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);`
	
	_, err := r.db.Exec(query)
	return err
}

// FindAll returns all tasks
func (r *TaskRepository) FindAll() ([]Task, error) {
	query := `SELECT id, title, description, completed, due_date, completed_at, created_at, updated_at FROM tasks ORDER BY created_at DESC`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var tasks []Task
	
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.DueDate, &t.CompletedAt, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	
	return tasks, nil
}

// FindByID returns a task by ID
func (r *TaskRepository) FindByID(id int) (*Task, error) {
	query := `SELECT id, title, description, completed, due_date, completed_at, created_at, updated_at FROM tasks WHERE id = ?`
	
	var t Task
	err := r.db.QueryRow(query, id).Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.DueDate, &t.CompletedAt, &t.CreatedAt, &t.UpdatedAt)
	
	if err == sql.ErrNoRows {
		return nil, errors.New("task not found")
	} else if err != nil {
		return nil, err
	}
	
	return &t, nil
}

// Create adds a new task
func (r *TaskRepository) Create(task *Task) (*Task, error) {
	query := `
	INSERT INTO tasks (title, description, completed, due_date, completed_at, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	RETURNING id`
	
	err := r.db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.DueDate,
		task.CompletedAt,
		task.CreatedAt,
		task.UpdatedAt,
	).Scan(&task.ID)
	
	if err != nil {
		return nil, err
	}
	
	return task, nil
}

// Update modifies an existing task
func (r *TaskRepository) Update(task *Task) (*Task, error) {
	query := `
	UPDATE tasks 
	SET title = ?, description = ?, completed = ?, due_date = ?, completed_at = ?, updated_at = ?
	WHERE id = ?`
	
	_, err := r.db.Exec(
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.DueDate,
		task.CompletedAt,
		task.UpdatedAt,
		task.ID,
	)
	
	if err != nil {
		return nil, err
	}
	
	return task, nil
}

// Delete removes a task
func (r *TaskRepository) Delete(id int) error {
	query := `DELETE FROM tasks WHERE id = ?`
	
	_, err := r.db.Exec(query, id)
	return err
}
