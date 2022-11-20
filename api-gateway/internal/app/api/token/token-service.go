package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	Id int `json:"id"`
}

const (
	salt       = "NGNWDYxigegc3yugIGCE&co33rffIOU&@*B@E("
	signingKey = "NGNWDYxigegc3yugIGCE&co33rffIOU&@*B@E("
	tokenTTL   = 12 * time.Hour
)

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		fmt.Println(signingKey)

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	fmt.Println(claims.Id)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Id, nil
}
