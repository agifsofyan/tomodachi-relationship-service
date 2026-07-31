package app

import (
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/config"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/database"
)

type Dependencies struct {
	Config      *config.Config
	Database    *database.Client
	Transaction database.Transaction

	// Logger logger.Logger // nanti kalau sudah ada abstraction
	// Validator *validator.Validator // nanti
}

func NewDependencies(
	cfg *config.Config,
	db *database.Client,
) *Dependencies {

	return &Dependencies{
		Config:      cfg,
		Database:    db,
		Transaction: database.NewTransaction(db.DB),
	}
}
