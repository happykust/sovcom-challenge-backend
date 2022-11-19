package auth

import (
	logger "account/pkg/logging"
	LoggerTypes "account/pkg/logging/types"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/matthewhartstonge/argon2"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	Id uint `json:"id"`
}

const (
	tokenTTL = 12 * time.Hour
)

func HashUserPassword(password string) string {
	passwordWithSecret := password + os.Getenv("SALT")
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(passwordWithSecret), bcrypt.DefaultCost)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Could not hash password", err)
	}
	return string(hashedBytes)
}

func validatePassword(password string, hashedPassword string) bool {
	passwordWithSecret := password + os.Getenv("SALT")
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordWithSecret))
	return err == nil
}

func GenerateTokens(id uint) (string, string) {

	AccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix()},
		id,
	})
	RefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})
	AccessTokenString, err := AccessToken.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Could not sign access token", err)
	}
	RefreshTokenString, err := RefreshToken.SignedString([]byte(os.Getenv("SIGNING_KEY")))
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Could not sign refresh token", err)
	}
	HashRefreshToken := HashRfToken(RefreshTokenString)
	return AccessTokenString, HashRefreshToken
}

func RefreshToken(id uint, rfToken string) (string, string) {
	if validateRfToken(id, rfToken) {
		return GenerateTokens(id)
	}
	return "", ""
}

func validateRfToken(id uint, rfToken string) bool {
	user := GetUnverifiedUserById(id)
	if user == nil || len(user[0].RefreshTokenHash) == 0 {
		fmt.Println("User not found")
		return false
	}
	rtMatches, err := argon2.Decode([]byte(user[0].RefreshTokenHash))
	ok, err := rtMatches.Verify([]byte(rfToken))
	fmt.Println(ok)
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Could not verify refresh token", err)
		return false
	}
	if ok == true {
		return true
	}
	return false
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Log(LoggerTypes.CRITICAL, "Could not parse token", nil)
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SIGNING_KEY")), nil
	})
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Could not parse token", err)
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	fmt.Println(claims.Id)
	if !ok {
		logger.Log(LoggerTypes.CRITICAL, "Could not parse token", err)
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return int(claims.Id), nil
}

func HashRfToken(token string) string {
	cfg := argon2.DefaultConfig()
	hash, err := cfg.Hash([]byte(token), []byte(os.Getenv("SALT")))
	encoded := hash.Encode()
	if err != nil {
		logger.Log(LoggerTypes.CRITICAL, "Could not hash refresh token", err)
	}
	return string(encoded)
}
