package token

import (
	"api-gateway/internal/app/api/domain/account/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func validateRfToken(c *gin.Context) {

	rfToken, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token"})

	}
	if rfToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token"})
		return
	}

	userId, err := ParseToken(rfToken)
	if err != nil {
		auth.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
	fmt.Print(userId)

}

func validateAccToken(c *gin.Context) {

	accToken, err := c.Cookie("access_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing access token"})
		return
	}
	if accToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing access token"})
		return
	}

	userId, err := ParseToken(accToken)
	if err != nil {
		auth.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
	fmt.Print(userId)

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
