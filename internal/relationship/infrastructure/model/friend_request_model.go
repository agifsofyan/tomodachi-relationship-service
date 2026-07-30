package model

import (
	"time"

	enum "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/enums"
	"github.com/google/uuid"
)

type FriendRequestModel struct {
	ID          uuid.UUID                `gorm:"type:uuid;primaryKey"`
	RequesterID uuid.UUID                `gorm:"type:uuid;not null;index"`
	ReceiverID  uuid.UUID                `gorm:"type:uuid;not null;index"`
	Status      enum.FriendRequestStatus `gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (FriendRequestModel) TableName() string {
	return "friend_requests"
}
