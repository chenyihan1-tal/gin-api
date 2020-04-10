package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(code int, msg string, data interface{}) {
	if msg == "" { msg = http.StatusText(code)}
	if data == nil { data = gin.H{} }

	g.C.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	return
}
