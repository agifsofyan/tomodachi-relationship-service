package sendfriendrequest

import (
	"net/http"

	sendfriendrequest "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/application/send_friend_request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SendFriendHandler struct {
	sendFriendService *sendfriendrequest.SendFriendService
}

func NewFriendRequestHandler(
	sendFriendService *sendfriendrequest.SendFriendService,
) *SendFriendHandler {

	return &SendFriendHandler{
		sendFriendService: sendFriendService,
	}
}

func (h *SendFriendHandler) Handle(
	c *gin.Context,
) {

	var request SendFriendRequestHttp

	if err := c.ShouldBindJSON(&request); err != nil {
		// validator response
		return
	}

	userID := uuid.New()

	response, err := h.sendFriendService.Execute(
		c.Request.Context(),
		ToSendFriendApplicationRequest(
			userID,
			request,
		),
	)
	if err != nil {
		// error mapper
		return
	}

	c.JSON(
		http.StatusCreated,
		ToSendFriendResponse(response),
	)
}
