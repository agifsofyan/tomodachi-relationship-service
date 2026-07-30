package auth

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	jwt.RegisteredClaims
}
