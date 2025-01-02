package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/krossroad/imperialfleet/internal"
	"github.com/krossroad/imperialfleet/internal/logger"
)

var log *logger.Entry

func main() {
	log = bootstrapLogger()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	opt := internal.Options{
		DBUrl:    os.Getenv("DB_URL"),
		HTTPAddr: os.Getenv("HTTP_ADDR"),
	}

	svc, err := internal.NewService(ctx, log, opt)
	if err != nil {
		log.Fatal("failed to create service", "error", err)
	}

	shutdownOnSignal(ctx, svc)
}

const waitTime = 10 * time.Second

func shutdownOnSignal(ctx context.Context, svc *internal.Service) {
	sig := waitForSignal()
	log.With("shutdown_signal", sig).Info("shutting down gracefully")

	ctx, cancel := context.WithTimeout(ctx, waitTime)
	defer cancel()

	svc.Stop(ctx)
	log.Info("shutting down server")
	os.Exit(0)
}

func waitForSignal() string {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	sig := <-c

	return sig.String()
}

func bootstrapLogger() *logger.Entry {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	return logger.New(handler)
}
