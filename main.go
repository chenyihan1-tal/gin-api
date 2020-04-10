package main

import (
	"fmt"
	"log"
	"net/http"

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

// @title Golang Gin API
// @version 1.0
// @description An example of gin
func main() {
	gin.SetMode(conf.App.RunMode)

	routersInit := router.Setup()
	readTimeout := conf.App.ReadTimeout
	writeTimeout := conf.App.WriteTimeout
	endPoint := fmt.Sprintf(":%d", conf.App.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
