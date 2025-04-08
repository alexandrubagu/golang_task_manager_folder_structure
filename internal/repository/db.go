package repository

import (
	"database/sql"
	"strings"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// NewDatabase creates a new database connection
func NewDatabase(url string) (*sql.DB, error) {
	// Parse URL to get driver and data source
	parts := strings.SplitN(url, "://", 2)
	if len(parts) != 2 {
		return nil, ErrInvalidDatabaseURL
	}
	
	driver := parts[0]
	dataSource := parts[1]
	
	// Open database connection
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}
	
	// Test connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	// Initialize task repository
	taskRepo := NewTaskRepository(db)
	if err := taskRepo.Initialize(); err != nil {
		return nil, err
	}
	
	return db, nil
}

// Error definitions
var (
	ErrInvalidDatabaseURL = New("invalid database URL")
)

// New creates a new error
func New(text string) error {
	return error(errorString(text))
}

type errorString string

func (e errorString) Error() string {
	return string(e)
}
