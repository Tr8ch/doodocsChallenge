package main

import (
	"doodocs/config"
	archiveinfo "doodocs/internal/http-server/handlers/archiveInfo"
	"doodocs/internal/http-server/handlers/archiving"
	emailsender "doodocs/internal/http-server/handlers/emailSender"
	mwLogger "doodocs/internal/http-server/middleware/logger"
	"doodocs/internal/lib/logger/handlers/slogpretty"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting doodocs backend challenge", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	router := chi.NewRouter()

	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)

	router.Post("/api/archive/information", archiveinfo.New(log))
	router.Post("/api/archive/files", archiving.New(log))
	router.Post("/api/mail/file", emailsender.New(log))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
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

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
