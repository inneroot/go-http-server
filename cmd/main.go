package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/inneroot/go-http-server/internal/handlers"
	"github.com/inneroot/go-http-server/internal/server"
	"github.com/inneroot/go-http-server/pkg/logger"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log := logger.SetLogger()

	srv := server.New("8080", log)
	handlers.Initialize()
	srv.Start()

	<-ctx.Done() // graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server shutdown failed", slog.Any("error", err))
	}
	log.Info("graceful shutdown finished")
}
