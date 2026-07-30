package sendfriendrequest

import (
	app "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/application/send_friend_request"

	"github.com/google/uuid"
)

func ToSendFriendApplicationRequest(
	userID uuid.UUID,
	request SendFriendRequestHttp,
) app.SendFriendRequest {

	return app.SendFriendRequest{
		RequesterID: userID,
		ReceiverID:  request.ReceiverID,
	}
}

func ToSendFriendResponse(
	response *app.SendFriendResponse,
) *SendFriendResponseHttp {

	return &SendFriendResponseHttp{
		ID:          response.ID,
		RequesterID: response.RequesterID,
		ReceiverID:  response.ReceiverID,
		Status:      string(response.Status),
		CreatedAt:   response.CreatedAt,
	}
}
