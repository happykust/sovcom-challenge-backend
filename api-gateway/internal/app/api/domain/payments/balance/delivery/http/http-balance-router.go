package http

import (
	"api-gateway/internal/app/api/domain/payments/balance"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	balanceRoutes := route.Group("/balance")
	balanceRoutes.GET("/get", balance.GetBalance)
}
