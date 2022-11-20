package payments

import (
	balanceHttpRoutes "api-gateway/internal/app/api/domain/payments/balance/delivery/http"
	transactionHttpRoutes "api-gateway/internal/app/api/domain/payments/transaction/delivery/http"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	paymentsRoutes := route.Group("/payments")
	balanceHttpRoutes.Routes(paymentsRoutes)
	transactionHttpRoutes.Routes(paymentsRoutes)
}
