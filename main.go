package main

import (
	"fmt"
	"gin-server-api/models"
	"gin-server-api/pkg/logger"
	"gin-server-api/pkg/setting"
	"gin-server-api/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d",setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:endPoint,
		ReadTimeout:readTimeout,
		WriteTimeout:writeTimeout,
		MaxHeaderBytes:maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s",endPoint)
	server.ListenAndServe()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	_ = router.Run() // listen and serve on 0.0.0.0:8080
}
