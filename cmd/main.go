package main

import (
	"context"
	"errors"
	"github.com/RIDOS/news-blog/internal/app/config"
	"github.com/RIDOS/news-blog/internal/app/logger"
	"github.com/RIDOS/news-blog/pkg/handler"
	"github.com/RIDOS/news-blog/pkg/repository"
	"github.com/RIDOS/news-blog/storage"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	log.Info("Server start", slog.String("env", cfg.Env))

	st, err := storage.NewStorage(cfg.StorageType, cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", slog.Attr{
			Key:   "err",
			Value: slog.StringValue(err.Error()),
		})
		os.Exit(1)
	}

	log.Info("Storage start", slog.String("storage_path", cfg.StoragePath))

	newsRepo := repository.NewNewsRepository(st)
	newsHandler := handler.NewNewsHandler(log, newsRepo)
	newsHandler.SetupRoutes()

	server := &http.Server{
		Addr: ":8080",
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Info("Listening on :8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("failed to start server", slog.Attr{
				Key:   "err",
				Value: slog.StringValue(err.Error()),
			})
			os.Exit(1)
		}
	}()

	<-stop
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("failed to gracefully shutdown server", slog.Attr{
			Key:   "err",
			Value: slog.StringValue(err.Error()),
		})
	}

	log.Info("Server stopped")
}
