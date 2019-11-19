package main

import (
	"gin-server-api/models"
	"gin-server-api/pkg/logger"
	"gin-server-api/pkg/setting"
	"gin-server-api/pkg/util"
	"github.com/gin-gonic/gin"
)

func init(){
	setting.Setup()
	models.Setup()
	logger.Setup()
	util.Setup()
}

func main(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
