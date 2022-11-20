package http

import (
	"api-gateway/internal/app/api/domain/deals/currencyDeals"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	currencyDealsRoutes := route.Group("/currency-deals")
	currencyDealsRoutes.POST("/buy", currencyDeals.CreateCurrencyDeal)
	currencyDealsRoutes.POST("/delete", currencyDeals.DeleteCurrencyDeal)
	currencyDealsRoutes.GET("/read", currencyDeals.ReadCurrencyDeal)
}
