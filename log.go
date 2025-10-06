package main

import (
	"log/slog"
	"os"
	"strings"
)

const LogLevel = "LOG_LEVEL"

func parseLogLevel(lvl string) slog.Level {
	switch strings.ToLower(lvl) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func newLogger() (*slog.Logger, slog.Level) {
	logLevel := parseLogLevel(os.Getenv(LogLevel))
	options := slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	}
	th := slog.NewTextHandler(os.Stderr, &options)
	logger := slog.New(th)

	return logger, logLevel
}
