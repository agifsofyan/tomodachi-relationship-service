package logger

import (
	"log/slog"
	"os"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/config"
)

func New(cfg *config.Config) *slog.Logger {

	var handler slog.Handler

	opts := &slog.HandlerOptions{
		Level: parseLevel(cfg.Logger.Level),
	}

	if cfg.App.Env == "production" {

		handler = slog.NewJSONHandler(os.Stdout, opts)

	} else {

		handler = slog.NewTextHandler(os.Stdout, opts)

	}

	return slog.New(handler)
}
