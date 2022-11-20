package authHttpRouter

import (
	"api-gateway/internal/app/api/domain/account/auth"
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
