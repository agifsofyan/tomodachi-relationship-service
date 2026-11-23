package config

type Config struct {
	App      AppConfig      `yaml:"app"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JwtConfig      `yaml:"jwt"`
	Logger   LoggerConfig   `yaml:"logger"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Env     string `yaml:"env"`
}

type ServerConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	ReadTimeout  int    `yaml:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout"`
}

type DatabaseConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	Name            string `yaml:"name"`
	SSLMode         string `yaml:"sslMode"`
	Timezone        string `yaml:"timezone"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}

type JwtConfig struct {
	SecretKey            string `yaml:"secret_key"`
	Algorithm            string `yaml:"algorithm"`
	AccessTokenExpiryDay int    `yaml:"access_token_expiry_day"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}
