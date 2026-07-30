package sendfriendrequest

import "github.com/google/uuid"

type SendFriendRequest struct {
	RequesterID uuid.UUID
	ReceiverID  uuid.UUID
}
