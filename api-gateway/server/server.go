package server

import (
	"api-gateway/internal/app/api/router"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {

	appRouter := gin.Default()

	appRouter.SetTrustedProxies([]string{"*"})

	router.Routes(appRouter)

	return appRouter

}
