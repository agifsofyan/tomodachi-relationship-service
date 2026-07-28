package database

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg *config.Config, log *slog.Logger) (*Client, error) {

	db, err := gorm.Open(
		postgres.Open(cfg.Database.DSN()),
		&gorm.Config{},
	)

	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(
		time.Duration(cfg.Database.ConnMaxLifetime) * time.Second,
	)

	// Pool configuration...
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Info("database connected")

	return &Client{
		DB: db,
	}, nil
}
