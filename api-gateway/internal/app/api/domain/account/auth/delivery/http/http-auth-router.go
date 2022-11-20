package authHttpRouter

import (
	"api-gateway/internal/app/api/domain/account/auth"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	authRoutes := route.Group("/auth")
	authRoutes.POST("/local/sign-up", auth.CreateUserEvent)
	authRoutes.POST("/local/sign-in", auth.LoginUserEvent)
	authRoutes.POST("/local/refresh", auth.RefreshTokenEvent)
	authRoutes.POST("/local/sign-out", auth.LogoutUserEvent)
	return
}

const (
	userId       = "userId"
	userBan      = "userBan"
	userVerified = "userVerified"
	userRole     = "userRole"
)

func Test(c *gin.Context) {
	f, err := c.Get("userId")
	if err {
		c.JSON(500, err)
	}
	fmt.Println(f)

	c.JSON(200, f)
}
