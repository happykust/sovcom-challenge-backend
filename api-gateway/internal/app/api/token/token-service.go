package token

import (
	logger "api-gateway/pkg/logging"
	LoggerTypes "api-gateway/pkg/logging/types"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Role string

const (
	RoleDeveloper Role = "developer"
	RoleAdmin     Role = "admin"
	RoleUser      Role = "user"
)

type tokenClaims struct {
	jwt.StandardClaims
	Id           uint `json:"id"`
	UserVerified bool `json:"user_verified"`
	Role         Role `json:"role"`
	Ban          bool `json:"ban"`
}

const (
	tokenTTL = 12 * time.Hour
)

func ParseToken(accessToken string) (*tokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Log(LoggerTypes.CRITICAL, "Could not parse token", nil)
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Could not parse token", err)
		return nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	fmt.Println(claims.Id)
	if !ok {
		logger.Log(LoggerTypes.CRITICAL, "Could not parse token", err)
		return nil, errors.New("token claims are not of type *tokenClaims")
	}
	return claims, nil
}
