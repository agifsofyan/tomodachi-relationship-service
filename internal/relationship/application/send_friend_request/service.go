package sendfriendrequest

import (
	"context"
)

type SendFriendService struct {
	SendFriendServiceDependencies
}

func NewSendFriendService(
	deps SendFriendServiceDependencies,
) *SendFriendService {

	return &SendFriendService{
		SendFriendServiceDependencies: deps,
	}
}

func (s *SendFriendService) Execute(
	ctx context.Context,
	request SendFriendRequest,
) (*SendFriendResponse, error) {

	// Business validation
	if err := s.validate(ctx, request); err != nil {
		return nil, err
	}

	// Persist friend request
	friendRequest, err := s.createFriendRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	// Convert to response DTO
	return s.toResponse(friendRequest), nil
}
