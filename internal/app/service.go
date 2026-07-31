package app

import (
	application "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/application/send_friend_request"
)

type Services struct {
	SendFriend *application.SendFriendService
}

func NewServices(
	deps *Dependencies,
	repo *Repositories,
) *Services {

	return &Services{

		SendFriend: application.NewSendFriendService(
			application.SendFriendServiceDependencies{
				Transaction:             deps.Transaction,
				FriendRequestRepository: repo.FriendRequest,
				FriendshipRepository:    repo.Friendship,
				BlockRepository:         repo.Block,
			},
		),
	}
}
