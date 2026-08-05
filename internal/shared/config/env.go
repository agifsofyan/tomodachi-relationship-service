package config

func overrideEnv(cfg *Config) {
	cfg.App.Env = getenv("APP_ENV", cfg.App.Env)
	cfg.Server.Port = getenvInt("SERVER_PORT", cfg.Server.Port)
	cfg.Database.Host = getenv("DB_HOST", cfg.Database.Host)
	cfg.Database.Port = getenvInt("DB_PORT", cfg.Database.Port)
	cfg.Database.User = getenv("DB_USER", cfg.Database.User)
	cfg.Database.Password = getenv("DB_PASSWORD", cfg.Database.Password)
	cfg.Database.Name = getenv("DB_NAME", cfg.Database.Name)
	cfg.Logger.Level = getenv("LOGGER_LEVEL", cfg.Logger.Level)
}
