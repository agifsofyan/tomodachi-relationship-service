package entity

import "time"

type Block struct {
	ID string

	UserID        string
	BlockedUserID string

	CreatedAt time.Time
}
