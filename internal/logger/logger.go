package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Logger provides logging capabilities
type Logger struct {
	level  string
	logger *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger(level string) *Logger {
	return &Logger{
		level:  strings.ToLower(level),
		logger: log.New(os.Stdout, "", 0),
	}
}

// Info logs information messages
func (l *Logger) Info(message string, args ...interface{}) {
	l.log("INFO", message, args...)
}

// Error logs error messages
func (l *Logger) Error(message string, err error, args ...interface{}) {
	if err != nil {
		message = fmt.Sprintf("%s: %v", message, err)
	}
	l.log("ERROR", message, args...)
}

// Debug logs debug messages
func (l *Logger) Debug(message string, args ...interface{}) {
	if l.level == "debug" {
		l.log("DEBUG", message, args...)
	}
}

// Fatal logs fatal messages and exits the program
func (l *Logger) Fatal(message string, err error, args ...interface{}) {
	if err != nil {
		message = fmt.Sprintf("%s: %v", message, err)
	}
	l.log("FATAL", message, args...)
	os.Exit(1)
}

func (l *Logger) log(level, message string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}
	l.logger.Printf("[%s] %s: %s", level, timestamp, message)
}
