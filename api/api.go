package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResOK struct {
	Code int         `json:"code" example:"200"`
	Msg  string      `json:"msg" example:"OK"`
	Data interface{} `json:"data"`
}

type ResError struct {
	Code int         `json:"code" example:"400"`
	Msg  string      `json:"msg" example:"Bad Request"`
	Data interface{} `json:"data"`
}

type ResForbidden struct {
	Code int         `json:"code" example:"403"`
	Msg  string      `json:"msg" example:"Forbidden"`
	Data interface{} `json:"data"`
}

type ResNotFound struct {
	Code int         `json:"code" example:"404"`
	Msg  string      `json:"msg" example:"Not Found"`
	Data interface{} `json:"data"`
}

func (g *Gin) ResOK(data ...interface{}) {
	if len(data) <= 0 {
		data = append(data, gin.H{})
	}

	g.C.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
		Data: data[0],
	})
	g.C.Abort()

	return
}

func (g *Gin) ResError(msg string, code ...int) {
	if len(code) <= 0 {
		code = append(code, http.StatusBadRequest)
	}

	if msg == "" {
		msg = http.StatusText(code[0])
	}

	g.C.JSON(http.StatusOK, Response{
		Code: code[0],
		Msg:  msg,
		Data: gin.H{},
	})
	g.C.Abort()

	return
}

func (g *Gin) Response(code int, msg string, data interface{}) {
	if msg == "" {
		msg = http.StatusText(code)
	}

	if data == nil {
		data = gin.H{}
	}

	g.C.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	g.C.Abort()

	return
}
