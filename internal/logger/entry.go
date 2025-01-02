package logger

import (
	"log/slog"
	"syscall"
)

type Entry struct {
	*slog.Logger
}

func NewEntry(logger *slog.Logger) *Entry {
	return &Entry{Logger: logger}
}

func (e *Entry) Fatal(msg string, fields ...interface{}) {
	e.Logger.Error(msg, fields...)
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
}

func New(handler slog.Handler) *Entry {
	return NewEntry(slog.New(handler))
}
