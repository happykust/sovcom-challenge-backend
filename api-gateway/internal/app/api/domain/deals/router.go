package deals

import (
	currencyHttpRouter "api-gateway/internal/app/api/domain/deals/currencyDeals/delivery/http"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	dealsRoutes := route.Group("/deals")
	currencyHttpRouter.Routes(dealsRoutes)
}
