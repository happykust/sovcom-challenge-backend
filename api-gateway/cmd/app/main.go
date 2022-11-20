package main

import (
	"api-gateway/pkg/core/config"
	"api-gateway/pkg/s3"
	"api-gateway/server"
	"fmt"
)

func init() {
	config.Init()
}

func main() {
	s3.InitS3Connection()
	app := server.App()
	err := app.Run(":9090")
	if err != nil {
		fmt.Println(err)
		return
	}
	select {}
}
