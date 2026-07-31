package auth

import (
	"errors"

	domainError "github.com/agifsofyan/tomodachi-relationship-service/internal/relationship/domain/errors"

	"github.com/golang-jwt/jwt/v5"
)

var ErrUnexpectedSigningMethod = errors.New("unexpected signing method")

type TokenParser interface {
	Parse(token string) (*JwtClaims, error)
}

type JWTParser struct {
	secret    []byte
	algorithm string
}

func NewParser(
	secret string,
	algorithm string,
) TokenParser {
	return &JWTParser{
		secret:    []byte(secret),
		algorithm: algorithm,
	}
}

func (p *JWTParser) Parse(
	tokenString string,
) (*JwtClaims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&JwtClaims{},
		p.keyFunc,
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid {
		return nil, domainError.ErrUnauthorized
	}

	return claims, nil
}

func (p *JWTParser) keyFunc(
	token *jwt.Token,
) (any, error) {

	if token.Method.Alg() != p.algorithm {
		return nil, ErrUnexpectedSigningMethod
	}

	return p.secret, nil
}
