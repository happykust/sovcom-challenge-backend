package router

import (
	authHttpRouter "api-gateway/internal/app/api/domain/account/auth/delivery/http"
	dealsHttpRouter "api-gateway/internal/app/api/domain/deals"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	routesAPIGroup := route.Group("/api")

	authHttpRouter.Routes(routesAPIGroup)
	dealsHttpRouter.Routes(routesAPIGroup)
	return

}
