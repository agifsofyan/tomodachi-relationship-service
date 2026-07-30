package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/agifsofyan/tomodachi-relationship-service/internal/shared/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthMiddleware struct {
	parser auth.TokenParser
}

func NewAuthMiddleware(
	parser auth.TokenParser,
) *AuthMiddleware {

	return &AuthMiddleware{
		parser: parser,
	}
}

func (m *AuthMiddleware) Handler() gin.HandlerFunc {

	return func(c *gin.Context) {

		token, ok := bearerToken(c)
		if !ok {
			unauthorized(c)
			return
		}

		claims, err := m.parser.Parse(token)
		if err != nil {
			unauthorized(c)
			return
		}

		userID, err := uuid.Parse(claims.Subject)
		if err != nil {
			unauthorized(c)
			return
		}

		ctx := withAuthContext(
			c.Request.Context(),
			userID,
			claims,
		)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func withAuthContext(
	ctx context.Context,
	userID uuid.UUID,
	claims *auth.JwtClaims,
) context.Context {

	ctx = auth.WithUserID(ctx, userID)
	ctx = auth.WithClaims(ctx, claims)

	return ctx
}

func unauthorized(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
}

func bearerToken(c *gin.Context) (string, bool) {

	header := c.GetHeader("Authorization")

	if header == "" {
		return "", false
	}

	if !strings.HasPrefix(header, "Bearer ") {
		return "", false
	}

	return strings.TrimPrefix(header, "Bearer "), true
}
