package router

import (
	authHttpRouter "api-gateway/internal/app/api/domain/account/auth/delivery/http"
	dealsHttpRouter "api-gateway/internal/app/api/domain/deals"
	loyalityHttpRouter "api-gateway/internal/app/api/domain/loyality"
	paymentsHttpRouter "api-gateway/internal/app/api/domain/payments"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	routesAPIGroup := route.Group("/api")

	authHttpRouter.Routes(routesAPIGroup)
	dealsHttpRouter.Routes(routesAPIGroup)
	loyalityHttpRouter.Routes(routesAPIGroup)
	paymentsHttpRouter.Routes(routesAPIGroup)
	return

}
