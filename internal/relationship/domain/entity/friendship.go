package entity

import "time"

type Friendship struct {
	ID string

	UserID   string
	FriendID string

	CreatedAt time.Time
}
