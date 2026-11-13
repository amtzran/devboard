// Command api runs the devboard HTTP API.
package main

import (
	"log/slog"
	"os"

	"github.com/amtzran/devboard/internal/handler"
	"github.com/amtzran/devboard/internal/middleware"
	"github.com/amtzran/devboard/internal/server"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	srv := server.New(":8000")
	srv.Use(middleware.Logger(logger))

	healthHandler := handler.NewHealth()
	srv.RegisterRoute("GET /health", healthHandler)

	logger.Info("Listening on", slog.Int("port", 8000))

	if err := srv.Start(); err != nil {
		logger.Error("server failed", slog.Any("error", err))
	}
}
