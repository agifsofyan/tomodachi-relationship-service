package sendfriendrequest

import "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/entity"

func (s *SendFriendService) toResponse(
	friendRequest *entity.FriendRequest,
) *SendFriendResponse {

	return &SendFriendResponse{
		ID:          friendRequest.ID,
		RequesterID: friendRequest.RequesterID,
		ReceiverID:  friendRequest.ReceiverID,
		Status:      friendRequest.Status,
		CreatedAt:   friendRequest.CreatedAt,
	}
}
