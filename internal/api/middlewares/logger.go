package middlewares

import (
	"net/http"
	"time"

	"golang_task_manager_folder_structure/internal/logger"
)

// LoggerMiddleware logs HTTP requests
func LoggerMiddleware(l *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Capture response with a wrapper
			ww := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			// Process request
			next.ServeHTTP(ww, r)

			// Log request details
			duration := time.Since(start)
			l.Info("%s %s %d %s", r.Method, r.RequestURI, ww.statusCode, duration)
		})
	}
}

// responseWriter is a wrapper for http.ResponseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
