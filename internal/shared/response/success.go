package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Success",
		Data:    data,
	})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "Created",
		Data:    data,
	})
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func OKWithPagination(
	c *gin.Context,
	data any,
	meta Pagination,
) {
	c.JSON(200, Response{
		Success: true,
		Message: "Success",
		Data:    data,
		Meta:    meta,
	})
}
