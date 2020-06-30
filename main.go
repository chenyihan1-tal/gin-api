package main

import (
	"fmt"
	"log"

	"github.com/xaviercry/gin-api/conf"
	"github.com/xaviercry/gin-api/model"
	"github.com/xaviercry/gin-api/package/cache"
	"github.com/xaviercry/gin-api/router"

	"github.com/gin-gonic/gin"
)

func init() {
	conf.Setup()
	cache.Setup()
	model.Setup()
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
