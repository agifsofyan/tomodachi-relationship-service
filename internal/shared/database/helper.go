package database

import (
	"context"

	"gorm.io/gorm"
)

func WithTransaction(
	ctx context.Context,
	tx *gorm.DB,
) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func DB(
	ctx context.Context,
	db *gorm.DB,
) *gorm.DB {

	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok {
		return tx
	}

	return db
}
