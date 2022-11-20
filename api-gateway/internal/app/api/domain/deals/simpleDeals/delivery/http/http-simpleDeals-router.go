package http

import (
	"api-gateway/internal/app/api/domain/deals/simpleDeals"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	currencyDealsRoutes := route.Group("/simple-deals")
	currencyDealsRoutes.POST("/buy", simpleDeals.CreateSimpleDeal)
}
