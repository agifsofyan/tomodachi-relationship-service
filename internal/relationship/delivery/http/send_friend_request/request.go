package sendfriendrequest

import "github.com/google/uuid"

type SendFriendRequestHttp struct {
	ReceiverID uuid.UUID `json:"receiver_id" validate:"required"`
}
