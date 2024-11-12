package main

import (
	"github.com/RIDOS/news-blog/handler"
	"github.com/RIDOS/news-blog/internal/config"
	"github.com/RIDOS/news-blog/internal/lib/logger/sl"
	"log/slog"
	"net/http"
	"os"
)

const (
	local = "local"
	dev   = "dev"
	prod  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("Server start", slog.String("env", cfg.Env))

	//storage, err := sqlite.New(cfg.StoragePath)
	//if err != nil {
	//	log.Error("failed to init storage", sl.Err(err))
	//	os.Exit(1)
	//}
	//
	//log.Info("Storage start", slog.String("env", cfg.Env))
	//_ = storage

	handler.SetupRoutes()

	log.Info("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error("failed to start server", sl.Err(err))
		os.Exit(1)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case local:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case dev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case prod:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
