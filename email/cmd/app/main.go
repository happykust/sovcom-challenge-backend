package main

import (
	"email/internal/app/routers"
	"email/pkg/core/config"
)

func init() {
	config.Init()
}

func main() {
	go routers.MainAmqpRouter()
	select {}
}
