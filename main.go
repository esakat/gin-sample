package main

import (
	"github.com/esakat/gin-sample/app"
	"github.com/esakat/gin-sample/router"
)

func main() {
	app := &app.App{}
	app.Init()

	router := router.Init(app)

	router.Run(":8080")
}