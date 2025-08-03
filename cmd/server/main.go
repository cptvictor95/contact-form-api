package main

import (
	logger "contact-form-api/internal"
	"net/http"
	"time"
)

func main() {
	// Initialize the global logger
	logger.Init()
	log := logger.Get()
	// Port for the server
	port := ":8000"

	log.Info("🚀 Starting contact form API server",
		"port", port,
		"environment", "development",
		"version", "1.0.0",
	)

	// fmt.Printf("Server is running on port %s\n", port)
	// fmt.Println("Try visiting http://localhost:8000/health to check if the server is running")

	http.HandleFunc("/health", withLogging(healthHandler))

	log.Error("Server failed to start", http.ListenAndServe(port, nil))
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseWriter := &responseWriter{ResponseWriter: w, statusCode: 200}

		next(responseWriter, r)

		logger.Get().Request(
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			responseWriter.statusCode,
			time.Since(start),
		)
	}
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write a JSON response
	response := `{"status": "healthy", "message": "Server is running"}`
	w.Write([]byte(response))
}