package logger

import (
	"log/slog"
	"os"
)

type Logger = *slog.Logger

func New(logLevelFlag string) Logger {
	logLevel := getLogLvel(logLevelFlag)

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	logger.Info("logger setup successfull", "loglevel", logLevel)

	return logger
}

func getLogLvel(flag string) slog.Level {
	logType := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}
	logLevel, ok := logType[flag]

	if !ok {
		logLevel = slog.LevelInfo
	}

	return logLevel
}
