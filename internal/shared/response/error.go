package response

import "github.com/gin-gonic/gin"

type ErrorItem struct {
	Field  string `json:"field,omitempty"`
	Reason string `json:"reason,omitempty"`
}

func Error(
	c *gin.Context,
	status int,
	message string,
	err any,
) {
	c.JSON(status, Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

func BadRequest(c *gin.Context, err any) {
	Error(c, 400, "Bad Request", err)
}

func Unauthorized(c *gin.Context, err any) {
	Error(c, 401, "Unauthorized", err)
}

func Forbidden(c *gin.Context, err any) {
	Error(c, 403, "Forbidden", err)
}

func NotFound(c *gin.Context, err any) {
	Error(c, 404, "Not Found", err)
}

func Conflict(c *gin.Context, err any) {
	Error(c, 409, "Conflict", err)
}

func Internal(c *gin.Context) {
	Error(c, 500, "Internal Server Error", nil)
}
