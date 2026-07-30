package sendfriendrequest

import (
	"time"

	"github.com/google/uuid"

	enum "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/enums"
)

type SendFriendResponse struct {
	ID          uuid.UUID
	RequesterID uuid.UUID
	ReceiverID  uuid.UUID
	Status      enum.FriendRequestStatus
	CreatedAt   time.Time
}
