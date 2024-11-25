package main

import (
	"Ferst/internal/config"
	"Ferst/internal/lib/logger/sl"
	"Ferst/internal/storage/sqllite"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	//config - cleaner
	log := setupLogger(cfg.Env)

	log.Info("starting url-shorts", slog.String("env", cfg.Env))
	log.Debug("debug message")

	// TODO: init storage - SQLlite
	storage, err := sqllite.New(cfg.Storage)
	if err != nil {
		log.Error("failed to sqlite", sl.Err(err))
		os.Exit(1)
	}

	_ = storage
	// TODO: init router - chi,render

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
