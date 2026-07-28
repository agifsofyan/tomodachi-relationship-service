package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		requestID, _ := c.Get(RequestIDKey)

		log.Info(
			"http_request",
			slog.String("request_id", requestID.(string)),
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.Int("status", c.Writer.Status()),
			slog.String("client_ip", c.ClientIP()),
			slog.Duration("duration", time.Since(start)),
		)
	}
}
