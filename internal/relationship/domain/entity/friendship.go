package entity

import (
	"time"

	"github.com/google/uuid"
)

type Friendship struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FriendID  uuid.UUID
	CreatedAt time.Time
}

func NewFriendship(
	userID uuid.UUID,
	friendID uuid.UUID,
) *Friendship {

	return &Friendship{
		ID:        uuid.New(),
		UserID:    userID,
		FriendID:  friendID,
		CreatedAt: time.Now(),
	}
}
