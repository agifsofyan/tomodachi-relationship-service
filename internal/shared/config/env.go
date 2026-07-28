package config

import (
	"os"
	"strconv"
)

func overrideEnv(cfg *Config) {

	if v := os.Getenv("APP_ENV"); v != "" {
		cfg.App.Env = v
	}

	if v := os.Getenv("SERVER_PORT"); v != "" {
		if port, err := strconv.Atoi(v); err == nil {
			cfg.Server.Port = port
		}
	}

	if v := os.Getenv("DB_HOST"); v != "" {
		cfg.Database.Host = v
	}

	if v := os.Getenv("DB_PORT"); v != "" {
		if port, err := strconv.Atoi(v); err == nil {
			cfg.Database.Port = port
		}
	}

	if v := os.Getenv("DB_USER"); v != "" {
		cfg.Database.User = v
	}

	if v := os.Getenv("DB_PASSWORD"); v != "" {
		cfg.Database.Password = v
	}

	if v := os.Getenv("DB_NAME"); v != "" {
		cfg.Database.Name = v
	}

	if v := os.Getenv("LOGGER_LEVEL"); v != "" {
		cfg.Logger.Level = v
	}
}
