package routers

import "github.com/gin-gonic/gin"

// init the routing information
func InitRouter() *gin.Engine{
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return r
}