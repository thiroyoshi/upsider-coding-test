package logger

import (
	"log/slog"
	"os"
)

type Handler struct {
	slog.Handler
}

func NewHandler() *Handler {
	opt := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}
	return &Handler{
		Handler: slog.NewJSONHandler(os.Stderr, &opt),
	}
}
