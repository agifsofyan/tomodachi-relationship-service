package sendfriendrequest

import (
	"context"

	domainError "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/errors"
)

func (s *SendFriendService) validate(
	ctx context.Context,
	request SendFriendRequest,
) error {

	if err := s.validateReceiver(ctx, request); err != nil {
		return err
	}

	if err := s.validateFriendship(ctx, request); err != nil {
		return err
	}

	if err := s.validatePendingRequest(ctx, request); err != nil {
		return err
	}

	if err := s.validateAlreadyFriend(ctx, request); err != nil {
		return err
	}

	if err := s.validateBlocked(ctx, request); err != nil {
		return err
	}

	return nil
}

func (s *SendFriendService) validateReceiver(
	ctx context.Context,
	request SendFriendRequest,
) error {

	return nil
}

func (s *SendFriendService) validateFriendship(
	ctx context.Context,
	request SendFriendRequest,
) error {

	exists, err := s.FriendshipRepository.Exists(
		ctx,
		request.RequesterID,
		request.ReceiverID,
	)
	if err != nil {
		return err
	}

	if exists {
		return domainError.ErrAlreadyFriend
	}

	return nil
}

func (s *SendFriendService) validatePendingRequest(
	ctx context.Context,
	request SendFriendRequest,
) error {

	friendRequest, err := s.FriendRequestRepository.FindPendingBetweenUsers(
		ctx,
		request.RequesterID,
		request.ReceiverID,
	)

	if err != nil {
		return err
	}

	if friendRequest != nil {
		return domainError.ErrFriendRequestAlreadyExists
	}

	return nil
}

func (s *SendFriendService) validateAlreadyFriend(
	ctx context.Context,
	request SendFriendRequest,
) error {

	exists, err := s.FriendshipRepository.Exists(
		ctx,
		request.RequesterID,
		request.ReceiverID,
	)
	if err != nil {
		return err
	}

	if exists {
		return domainError.ErrAlreadyFriend
	}

	return nil
}

func (s *SendFriendService) validateBlocked(
	ctx context.Context,
	request SendFriendRequest,
) error {

	blocked, err := s.BlockRepository.IsBlockedBetweenUsers(
		ctx,
		request.RequesterID,
		request.ReceiverID,
	)
	if err != nil {
		return err
	}

	if blocked {
		return domainError.ErrUserBlocked
	}

	return nil
}
