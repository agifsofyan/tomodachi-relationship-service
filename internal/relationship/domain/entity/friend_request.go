package entity

import "time"

type FriendRequest struct {
	ID string

	RequesterID string
	ReceiverID  string

	Status string

	CreatedAt time.Time
	UpdatedAt time.Time
}
