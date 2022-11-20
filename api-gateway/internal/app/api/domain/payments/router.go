package payments

import (
	balanceHttpRoutes "api-gateway/internal/app/api/domain/payments/balance/delivery/http"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	paymentsRoutes := route.Group("/payments")
	balanceHttpRoutes.Routes(paymentsRoutes)
}
