package app

import (
	delivery "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/delivery/http/send_friend_request"
)

type Handlers struct {
	SendFriend *delivery.SendFriendHandler
}

func NewHandlers(
	services *Services,
) *Handlers {

	return &Handlers{
		SendFriend: delivery.NewSendFriendHandler(
			services.SendFriend,
		),
	}
}
