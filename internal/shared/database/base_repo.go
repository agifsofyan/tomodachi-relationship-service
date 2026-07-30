package database

import (
	"context"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(
	db *gorm.DB,
) Repository {

	return Repository{
		DB: db,
	}
}

func (r Repository) Database(
	ctx context.Context,
) *gorm.DB {

	return DB(ctx, r.DB)
}
