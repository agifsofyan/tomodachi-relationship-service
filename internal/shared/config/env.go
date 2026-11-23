package config

func overrideEnv(cfg *Config) {
	// App
	cfg.App.Name = getenv("APP_NAME", cfg.App.Name)
	cfg.App.Version = getenv("APP_VERSION", cfg.App.Version)
	cfg.App.Env = getenv("APP_ENV", cfg.App.Env)

	// Server
	cfg.Server.Host = getenv("SERVER_HOST", cfg.Server.Host)
	cfg.Server.Port = getenvInt("SERVER_PORT", cfg.Server.Port)
	cfg.Server.ReadTimeout = getenvInt("SERVER_READ_TIMEOUT", cfg.Server.ReadTimeout)
	cfg.Server.WriteTimeout = getenvInt("SERVER_WRITE_TIMEOUT", cfg.Server.WriteTimeout)

	// Database
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

	// JWT
	cfg.JWT.SecretKey = getenv("SECRET_KEY", cfg.JWT.SecretKey)
	cfg.JWT.Algorithm = getenv("JWT_ALGORITHM", cfg.JWT.Algorithm)
	cfg.JWT.AccessTokenExpiryDay = getenvInt(
		"JWT_ACCESS_TOKEN_EXPIRY_DAY",
		cfg.JWT.AccessTokenExpiryDay,
	)

	// Logger
	cfg.Logger.Level = getenv("LOGGER_LEVEL", cfg.Logger.Level)
}
