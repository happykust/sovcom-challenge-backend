package server

import (
	"currency/internal/app/api/router"
	logger "currency/pkg/logging"
	LoggerTypes "currency/pkg/logging/types"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	appRouter := gin.Default()

	err := appRouter.SetTrustedProxies([]string{"*"})
	if err != nil {
		logger.Log(LoggerTypes.ERROR, "[Currency | API Server] Error set trusted proxies", err)
	}

	router.Routes(appRouter)
	return appRouter
}
