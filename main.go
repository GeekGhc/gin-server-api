package main

import (
	"gin-server-api/models"
	"gin-server-api/pkg/logger"
	"gin-server-api/pkg/setting"
	"gin-server-api/pkg/util"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	models.Setup()
	logger.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	gin.ForceConsoleColor()
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	_ = router.Run() // listen and serve on 0.0.0.0:8080
}
