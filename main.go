package main

import (
	"github.com/gin-gonic/gin"
	"tiktok_backend/config"
	"tiktok_backend/service"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)
	config.Connect()

	// Code below is for testing the database connection
	//config.InitTables()
	//config.InitFakeData()
	//config.CheckFakeData()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// The code below is mainly testing the installation for hertz
	//h := server.Default()
	//
	//h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
	//	ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	//})
	//
	//h.Spin()
	//db:=config.GetDB()

}
