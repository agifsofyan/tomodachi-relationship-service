package sendfriendrequest

import (
	"context"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"
)

func (s *SendFriendService) createFriendRequest(
	ctx context.Context,
	request SendFriendRequest,
) (*entity.FriendRequest, error) {

	friendRequest, err := entity.NewFriendRequest(
		request.RequesterID,
		request.ReceiverID,
	)
	if err != nil {
		return nil, err
	}

	err = s.Transaction.Execute(ctx, func(txCtx context.Context) error {
		return s.FriendRequestRepository.Create(txCtx, friendRequest)
	})
	if err != nil {
		return nil, err
	}

	return friendRequest, nil
}
