package token

import (
	"api-gateway/internal/app/api/domain/account/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	userId       = "userId"
	userBan      = "userBan"
	userVerified = "userVerified"
	userRole     = "userRole"
)

//
//func Middleware(c *gin.Context) {
//	//validateRfToken(c)
//	validateAccToken(c)
//}

//func validateRfToken(c *gin.Context) {
//
//	rfToken, err := c.Cookie("token")
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token"})
//
//	}
//	if rfToken == "" {
//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token"})
//		return
//	}
//
//	payloadR, err := ParseToken(rfToken)
//	if err != nil {
//		auth.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
//		return
//	}
//
//	c.Set(userCtx, payloadR)
//	fmt.Print(payloadR)
//
//}

func ValidateAccToken(c *gin.Context) *tokenClaims {

	accToken, err := c.Cookie("access_token")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing access token"})
		return nil
	}
	if accToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing access token"})
		return nil
	}

	payload, err := ParseToken(accToken)
	if err != nil {
		auth.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return nil
	}

	payload.Id = 6
	payload.UserVerified = true
	payload.Ban = false

	fmt.Println(payload)
	return payload

	//c.Set(userId, payload.Id)
	//c.Set(userBan, payload.Ban)
	//c.Set(userVerified, payload.UserVerified)
	//c.Set(userRole, payload.Role)
	////fmt.Println(payload.Id)
	//c.Next()
	//return nil
}

//func UserIdentify(c *gin.Context) {
//	header := c.GetHeader(authorizationHeader)
//	if header == "" {
//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
//		return
//	}
//
//	headerParts := strings.Split(header, " ")
//
//	if len(headerParts) != 2 {
//		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
//		return
//	}
//
//	if len(headerParts[1]) == 0 {
//		balance.NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
//		return
//	}
//
//	userId, err := ParseToken(headerParts[1])
//	fmt.Print(userId)
//	if err != nil {
//		balance.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
//		return
//	}
//
//	c.Set(userCtx, userId)
//
//}
//func GetUserId(c *gin.Context) (int, error) {
//	id, ok := c.Get(userCtx)
//	if !ok {
//		return 0, errors.New("transactions id not found")
//	}
//
//	idInt, ok := id.(int)
//	if !ok {
//		return 0, errors.New("transactions id is of invalid type")
//	}
//
//	return idInt, nil
//}
