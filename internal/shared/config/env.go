package config

func overrideEnv(cfg *Config) {
	cfg.App.Env = getenv("APP_ENV", cfg.App.Env)
	cfg.Server.Port = getenvInt("SERVER_PORT", cfg.Server.Port)

	cfg.Database.Host = getenv("DATABASE_HOST", cfg.Database.Host)
	cfg.Database.Port = getenvInt("DATABASE_PORT", cfg.Database.Port)
	cfg.Database.User = getenv("DATABASE_USER", cfg.Database.User)
	cfg.Database.Password = getenv("DATABASE_PASSWORD", cfg.Database.Password)
	cfg.Database.Name = getenv("DATABASE_NAME", cfg.Database.Name)
	cfg.Database.SSLMode = getenv("DATABASE_SSL_MODE", cfg.Database.SSLMode)
	cfg.Database.Timezone = getenv("DATABASE_TIMEZONE", cfg.Database.Timezone)
	cfg.Database.MaxOpenConns = getenvInt("DATABASE_MAX_OPEN_CONNS", cfg.Database.MaxOpenConns)
	cfg.Database.MaxIdleConns = getenvInt("DATABASE_MAX_IDLE_CONNS", cfg.Database.MaxIdleConns)
	cfg.Database.ConnMaxLifetime = getenvInt("DATABASE_CONN_MAX_LIFETIME", cfg.Database.ConnMaxLifetime)

	cfg.Logger.Level = getenv("LOGGER_LEVEL", cfg.Logger.Level)
}
