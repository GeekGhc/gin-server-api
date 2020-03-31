package app

import (
	"gin-server-api/helper/e"
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

type Message struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	if data != nil {
		g.C.JSON(httpCode, Response{
			Code: errCode,
			Msg:  e.GetMsg(errCode),
			Data: data,
		})
	} else {
		g.C.JSON(httpCode, Message{
			Code: errCode,
			Msg:  e.GetMsg(errCode),
		})
	}
	return
}
