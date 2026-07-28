package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery(log *slog.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		log.Error(
			"panic recovered",
			slog.Any("error", recovered),
		)

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal server error",
		})
	})
}
