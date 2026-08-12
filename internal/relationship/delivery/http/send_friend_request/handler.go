package sendfriendrequest

import (
	"net/http"

	sendfriendrequest "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/application/send_friend_request"
	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/auth"
	"github.com/gin-gonic/gin"
)

type SendFriendHandler struct {
	sendFriendService *sendfriendrequest.SendFriendService
}

func NewSendFriendHandler(
	sendFriendService *sendfriendrequest.SendFriendService,
) *SendFriendHandler {

	return &SendFriendHandler{
		sendFriendService: sendFriendService,
	}
}

// Handle godoc
// @Summary      Send friend request
// @Description  Send a friend request to another user
// @Tags         relationships
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body SendFriendRequestHttp true "Friend request payload"
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Router       /api/v1/relationships/friends/request [post]
func (h *SendFriendHandler) Handle(
	c *gin.Context,
) {

	var request SendFriendRequestHttp

	if err := c.ShouldBindJSON(&request); err != nil {
		// TODO: validator response
		return
	}

	userID, err := auth.UserID(c.Request.Context())
	if err != nil {
		// TODO: unauthorized response
		return
	}

	result, err := h.sendFriendService.Execute(
		c.Request.Context(),
		ToSendFriendApplicationRequest(
			userID,
			request,
		),
	)

	if err != nil {
		// TODO: error response
		return
	}

	c.JSON(
		http.StatusCreated,
		ToSendFriendResponse(result),
	)
}
