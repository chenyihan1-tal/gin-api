package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/xaviercry/gin-api/api"
	"github.com/xaviercry/gin-api/conf"
	_ "github.com/xaviercry/gin-api/docs"
)

// Setup initialize routing information
func Setup() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors())

	if conf.App.RunMode != "release" {
		r.Use(logs())
	}

	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/users", api.UserGetList) // 获取用户列表

	return r
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Set("content-type", "application/json")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func logs() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			return
		}

		start := time.Now()

		c.Next()

		end := time.Now()
		latency := end.Sub(start).Milliseconds()

		body, _ := ioutil.ReadAll(c.Request.Body)

		log.Printf("\n[%s][%d][%dms][%s][%s][%s][%s]",
			c.Request.Method,
			c.Writer.Status(),
			latency,
			c.Request.URL,
			string(body),
			c.Request.Host,
			c.Request.Referer(),
		)
	}
}
