package config

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JwtConfig      `mapstructure:"jwt"`
	Logger   LoggerConfig   `mapstructure:"logger"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Env     string `mapstructure:"env"`
}

type ServerConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"readTimeout"`
	WriteTimeout int    `mapstructure:"writeTimeout"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"sslMode"`
	Timezone string `yaml:"timezone"`

	MaxOpenConns    int `yaml:"maxOpenConns"`
	MaxIdleConns    int `yaml:"maxIdleConns"`
	ConnMaxLifetime int `yaml:"connMaxLifetime"`
}

type JwtConfig struct {
	SecretKey            string `yaml:"secret_key"`
	Algorithm            string `yaml:"algorithm"`
	AccessTokenExpiryDay int    `yaml:"access_token_expiry_day"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}
