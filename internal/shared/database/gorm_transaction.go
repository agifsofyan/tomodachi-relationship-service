package database

import (
	"context"

	"gorm.io/gorm"
)

type GormTransaction struct {
	db *gorm.DB
}

func NewTransaction(
	db *gorm.DB,
) Transaction {

	return &GormTransaction{
		db: db,
	}
}

func (t *GormTransaction) Execute(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {

	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		ctx = WithTransaction(ctx, tx)

		return fn(ctx)

	})
}
