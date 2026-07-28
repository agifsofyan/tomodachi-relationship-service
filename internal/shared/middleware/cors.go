package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
			"X-Request-ID",
		},
		ExposeHeaders: []string{
			"X-Request-ID",
		},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
}
