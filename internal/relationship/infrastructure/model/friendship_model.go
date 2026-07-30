package model

import (
	"time"

	"github.com/google/uuid"
)

type FriendshipModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`
	FriendID  uuid.UUID `gorm:"type:uuid;not null;index"`
	CreatedAt time.Time
}

func (FriendshipModel) TableName() string {
	return "friendships"
}
