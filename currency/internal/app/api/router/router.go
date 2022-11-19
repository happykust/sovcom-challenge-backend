package router

import (
	currency_ws "currency/internal/app/api/domain/currency-ws"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	routesAPIGroup := route.Group("/api")
	routesAPIGroup.GET("/websockets", func(c *gin.Context) {
		currency_ws.WebsocketHandler(c.Writer, c.Request)
	})
	return
}
