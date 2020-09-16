package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// NewConsoleLogger creates a new zerolog console logger
func NewConsoleLogger() zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	return zerolog.New(output).With().Timestamp().Logger()
}
