package logging

import (
	"log/slog"
	"os"
	"strings"
)

// NewLogger creates and returns a new slog logger.
// It parses the logLevel string to set the appropriate logging level.
func NewLogger(logLevel string) *slog.Logger {
	var level slog.Level
	switch strings.ToLower(logLevel) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo // Default to info level
	}

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	return slog.New(handler)
}
