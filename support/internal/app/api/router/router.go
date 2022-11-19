package router

import (
	"github.com/gin-gonic/gin"
	"support/internal/app/api/domain/support"
)

func Routes(route *gin.Engine) {
	routesAPIGroup := route.Group("/api")
	routesAPIGroup.GET("/websockets", func(c *gin.Context) {
		support.WebsocketHandler(c.Writer, c.Request)
	})
	return
}
