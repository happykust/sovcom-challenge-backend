package loyality

import (
	promocodesHttpRouter "api-gateway/internal/app/api/domain/loyality/promocodes/delivery/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	loyalityRoutes := route.Group("/loyality")
	promocodesHttpRouter.Routes(loyalityRoutes)
}
