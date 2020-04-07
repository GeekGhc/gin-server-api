package main

import (
	"flag"
	"gin-server-api/server"
)

func main() {
	//应用程序运行环境
	env := flag.String("env", "dev", "runtime env [dev|pre|prd]")
	flag.Parse()
	//启动程序 优雅结束
	server.NewGracefulServer().Run(*env).Wait()
}
