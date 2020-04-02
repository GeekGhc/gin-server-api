package main

import (
	"flag"
	"gin-server-api/app/setting"
	"gin-server-api/helper"
	"gin-server-api/internal/gredis"
	"gin-server-api/internal/logger"
	"gin-server-api/models"
	"gin-server-api/server"
	"gin-server-api/service/kafka_service"
)

func init() {
	setting.Setup()
	models.Setup()
	logger.Setup()
	gredis.Setup()
	helper.Setup()
	kafka_service.SetUp()
}

func main() {
	//应用程序运行环境
	env := flag.String("env", "dev", "runtime env [dev|pre|prd]")
	flag.Parse()
	//启动程序 优雅结束
	server.NewGracefulServer().Run(*env).Wait()
}
