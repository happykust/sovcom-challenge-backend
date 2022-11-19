package server

import (
	"github.com/gin-gonic/gin"
	"support/internal/app/api/router"
)

func App() *gin.Engine {

	appRouter := gin.Default()

	err := appRouter.SetTrustedProxies([]string{"*"})
	if err != nil {
		//logger
	}

	router.Routes(appRouter)

	return appRouter

}
