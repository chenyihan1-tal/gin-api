package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/xavierror/gowheel/tool"

	"github.com/xavierror/gin-api/api"
	_ "github.com/xavierror/gin-api/docs"
)

// Setup initialize routing information
func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logs())
	r.Use(cors())
	r.Use(gin.Recovery())

	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
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

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

type routerCache struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w routerCache) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

type reqLog struct {
	RequestID string              `json:"request_id"`
	Method    string              `json:"method"`
	Uri       string              `json:"uri"`
	Path      string              `json:"path"`
	Host      string              `json:"host"`
	Ip        string              `json:"ip"`
	Query     []string            `json:"query"`
	Header    map[string][]string `json:"header"`
	Body      string              `json:"body"`
	Response  string              `json:"response"`
	ReqTime   int64               `json:"req_time"`
	Since     int64               `json:"since"`
}

func logs() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			return
		}

		t := time.Now()

		requestBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		log := reqLog{
			Method:    c.Request.Method,
			Uri:       c.Request.RequestURI,
			Path:      c.Request.URL.Path,
			Host:      c.Request.Host,
			Query:     strings.Split(c.Request.URL.RawQuery, "&"),
			Ip:        c.ClientIP(),
			Header:    c.Request.Header,
			Body:      string(requestBody),
			ReqTime:   time.Now().Unix(),
			RequestID: tool.RandStr(32),
		}

		c.Header("x-request-id", log.RequestID)

		blw := &routerCache{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		log.Response = string(blw.body.Bytes())
		log.Since = time.Since(t).Milliseconds()

		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(log)

		// open, err := file.MustOpen("access.log", "runtime")
		// if err == nil {
		// 	logs.Error(err.Error())
		// } else {
		// 	open.Write(bf.Bytes())
		// 	open.Close()
		// }

		fmt.Println(bf.String())
	}
}
