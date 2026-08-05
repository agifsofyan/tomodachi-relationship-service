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
	var c Config

	file, err := os.ReadFile("configs/app.yml")
	if err == nil {
		if err := yaml.Unmarshal(file, &c); err != nil {
			return nil, fmt.Errorf("parse config: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("read config: %w", err)
	}

	overrideEnv(&c)

	return &c, nil
}
