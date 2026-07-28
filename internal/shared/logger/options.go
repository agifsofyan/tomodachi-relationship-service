package logger

import "log/slog"

func parseLevel(level string) slog.Level {

	switch level {

	case "debug":
		return slog.LevelDebug

	case "warn":
		return slog.LevelWarn

	case "error":
		return slog.LevelError

	default:
		return slog.LevelInfo
	}
}
