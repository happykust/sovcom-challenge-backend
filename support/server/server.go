package server

import (
	"github.com/gin-gonic/gin"
	"support/internal/app/api/router"
	logger "support/pkg/logging"
	LoggerTypes "support/pkg/logging/types"
)

func App() *gin.Engine {

	appRouter := gin.Default()

	err := appRouter.SetTrustedProxies([]string{"*"})
	if err != nil {
		logger.Log(LoggerTypes.ERROR, "Error while setting trusted proxies: ", err)
		panic(err)
	}

	router.Routes(appRouter)

	return appRouter

}
