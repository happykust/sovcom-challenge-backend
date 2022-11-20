package http

import (
	"api-gateway/internal/app/api/domain/loyality/promocodes"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	promocodesRoutes := route.Group("/promocodes")
	promocodesRoutes.POST("/create", promocodes.CreatePromocode)
	// promocodesRoutes.POST("/update", promocodes.UpdatePromocode)
	promocodesRoutes.POST("/delete", promocodes.DeletePromocode)
	promocodesRoutes.POST("/use", promocodes.UsePromocode)
	promocodesRoutes.GET("/read", promocodes.ReadPromocode)
}
