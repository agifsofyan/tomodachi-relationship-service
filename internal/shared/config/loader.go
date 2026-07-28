package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	cfg  *Config
	once sync.Once
)

func Load() (*Config, error) {
	var err error

	once.Do(func() {
		cfg, err = loadConfig()
	})

	return cfg, err
}

func loadConfig() (*Config, error) {

	file, err := os.ReadFile("configs/app.yml")
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var c Config

	if err := yaml.Unmarshal(file, &c); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	overrideEnv(&c)

	return &c, nil
}
