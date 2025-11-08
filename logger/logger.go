package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

var LOG *slog.Logger

func SetLogLevel() {
	logLevel := slog.LevelInfo
	level, found := os.LookupEnv("LOGLEVEL")

	if !found {
		fmt.Println("LOGLEVEL environment variable not set, defaulting to loglevel INFO")
		LOG = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
		return
	}

	switch strings.ToLower(level) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	LOG = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))

}
