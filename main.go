package main

import (
	"fmt"
	"gin-server-api/models"
	"gin-server-api/pkg/gredis"
	"gin-server-api/pkg/logger"
	"gin-server-api/pkg/setting"
	"gin-server-api/pkg/util"
	"gin-server-api/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func init() {
	setting.Setup()
	models.Setup()
	logger.Setup()
	gredis.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	gin.ForceConsoleColor()

	router := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout * time.Second,
		WriteTimeout:   writeTimeout * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.ListenAndServe()
}
