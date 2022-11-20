package deals

import (
	currencyHttpRouter "api-gateway/internal/app/api/domain/deals/currencyDeals/delivery/http"
	simpleHttpRouter "api-gateway/internal/app/api/domain/deals/simpleDeals/delivery/http"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	dealsRoutes := route.Group("/deals")
	currencyHttpRouter.Routes(dealsRoutes)
	simpleHttpRouter.Routes(dealsRoutes)
}
