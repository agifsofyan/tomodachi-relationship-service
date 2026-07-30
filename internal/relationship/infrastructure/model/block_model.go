package model

import (
	"time"

	"github.com/google/uuid"
)

type BlockModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	BlockerID uuid.UUID `gorm:"type:uuid;not null;index"`
	BlockedID uuid.UUID `gorm:"type:uuid;not null;index"`
	CreatedAt time.Time
}

func (BlockModel) TableName() string {
	return "blocks"
}
