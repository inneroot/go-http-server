package logger

import (
	"fmt"
	"log/slog"
	"os"
)

func SetLogger() *slog.Logger {
	level := slog.LevelInfo
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)
	logger.Info(fmt.Sprintf("Logger level: %s", level.String()), "ENV", os.Getenv("ENV"))
	return logger
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
