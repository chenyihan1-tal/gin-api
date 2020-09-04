package main

import (
	"fmt"
	"log"

	"github.com/xavierror/gin-api/conf"
	"github.com/xavierror/gin-api/model"
	"github.com/xavierror/gin-api/package/cache"
	"github.com/xavierror/gin-api/package/redLock"
	"github.com/xavierror/gin-api/package/single"
	"github.com/xavierror/gin-api/router"

	"github.com/gin-gonic/gin"
)

func init() {
	conf.Setup()
	cache.Setup()
	model.Setup()
	single.Setup()
	go redLock.Setup()
}

// @title GIN-API
// @version 1.0
// @description 响应码说明: 2xx请求成功，3xx重定向，4xx请求错误，5xx服务器错误
func main() {
	gin.SetMode(conf.App.RunMode)

	app := router.Setup()

	endPoint := fmt.Sprintf(":%d", conf.App.HttpPort)

	log.Printf("[info] start http server listening %s", endPoint)

	app.Run(endPoint)
}
