package main

import (
	"github.com/RIDOS/news-blog/internal/app/config"
	"github.com/RIDOS/news-blog/internal/app/logger"
	"github.com/RIDOS/news-blog/pkg/handler"
	"github.com/RIDOS/news-blog/pkg/repository"
	"github.com/RIDOS/news-blog/storage"
	"log/slog"
	"net/http"
	"os"
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

	log.Info("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error("failed to start server", slog.Attr{
			Key:   "err",
			Value: slog.StringValue(err.Error()),
		})
		os.Exit(1)
	}
}
