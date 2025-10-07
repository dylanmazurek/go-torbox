package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

func New(ctx context.Context) zerolog.Logger {
	envLogLevel := os.Getenv("LOG_LEVEL")
	if envLogLevel == "" {
		envLogLevel = zerolog.LevelInfoValue
	}

	logLevel, err := zerolog.ParseLevel(envLogLevel)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}

	writer := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05 01-02"}

	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.
		New(writer).
		With().
		Timestamp().
		Logger()

	return logger
}
