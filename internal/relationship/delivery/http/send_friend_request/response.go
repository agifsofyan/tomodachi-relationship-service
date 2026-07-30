package sendfriendrequest

import (
	"time"

	"github.com/google/uuid"
)

type SendFriendResponseHttp struct {
	ID          uuid.UUID `json:"id"`
	RequesterID uuid.UUID `json:"requester_id"`
	ReceiverID  uuid.UUID `json:"receiver_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
