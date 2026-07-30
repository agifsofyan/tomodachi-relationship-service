package entity

import (
	"time"

	enum "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/enums"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/errors"
	"github.com/google/uuid"
)

type FriendRequest struct {
	ID          uuid.UUID
	RequesterID uuid.UUID
	ReceiverID  uuid.UUID
	Status      enum.FriendRequestStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewFriendRequest(
	requesterID uuid.UUID,
	receiverID uuid.UUID,
) (*FriendRequest, error) {

	if requesterID == receiverID {
		return nil, errors.ErrCannotRequestSelf
	}

	now := time.Now()

	return &FriendRequest{
		ID:          uuid.New(),
		RequesterID: requesterID,
		ReceiverID:  receiverID,
		Status:      enum.FriendRequestPending,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (f *FriendRequest) Accept() {

	f.Status = enum.FriendRequestAccepted
	f.UpdatedAt = time.Now()

}

func (f *FriendRequest) Reject() {

	f.Status = enum.FriendRequestRejected
	f.UpdatedAt = time.Now()

}

func (f *FriendRequest) Cancel() {

	f.Status = enum.FriendRequestCancelled
	f.UpdatedAt = time.Now()

}
