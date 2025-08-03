package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"
)

type Logger struct {
	*slog.Logger
}

// Global logger instance
var Global *Logger

// New creates a new logger with a JSON handler and a timestamp attribute
func New() *Logger {
	// Create a custom pretty handler
	handler := &prettyHandler{
		handler: slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}),
	}
	
	logger := slog.New(handler)
	return &Logger{logger}
}

// prettyHandler formats logs in a more readable way
type prettyHandler struct {
	handler slog.Handler
}

func (h *prettyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *prettyHandler) Handle(ctx context.Context, r slog.Record) error {
	// Format the log entry nicely
	timestamp := time.Now().Format("15:04:05")
	level := r.Level.String()
	message := r.Message
	
	fmt.Printf("\n📝 [%s] %s: %s\n", timestamp, level, message)
	
	// Print each field on a new line
	r.Attrs(func(attr slog.Attr) bool {
		fmt.Printf("   %s: %v\n", attr.Key, attr.Value)
		return true
	})
	
	return nil
}

func (h *prettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &prettyHandler{handler: h.handler.WithAttrs(attrs)}
}

func (h *prettyHandler) WithGroup(name string) slog.Handler {
	return &prettyHandler{handler: h.handler.WithGroup(name)}
}

// Initialize the global logger
func Init() {
	Global = New()
}

// Get returns the global logger instance
func Get() *Logger {
	if Global == nil {
		Init()
	}

	return Global
}

// Info logs an info message with the given fields
func (l *Logger) Info(msg string, fields ...any) {
	l.Logger.Info(msg, fields...)
}

// Error logs an error message with the optional fields
func (l *Logger) Error(msg string, err error, fields ...any) {
  if err != nil {
		fields = append(fields, "error", err.Error())
	}
	l.Logger.Error(msg, fields...)
}

// Debug logs a debug message with optional fields
func (l *Logger) Debug(msg string, fields ...any) {
	l.Logger.Debug(msg, fields...)
}

func (l *Logger) Request(method, path, remoteAddr string, statusCode int, duration time.Duration) {
	// Color-coded status codes
	statusEmoji := "✅" // Green for 2xx
	if statusCode >= 400 {
		statusEmoji = "❌" // Red for 4xx/5xx
	} else if statusCode >= 300 {
		statusEmoji = "⚠️" // Yellow for 3xx
	}
	
	l.Info("HTTP Request",
		"method", method,
		"path", path,
		"status", statusEmoji,
		"code", statusCode,
		"duration", duration.String(),
		"ip", remoteAddr,
	)
}