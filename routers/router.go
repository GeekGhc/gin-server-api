package routers

import (
	"gin-server-api/middleware/jwt"
	"github.com/gin-gonic/gin"
)

// init the routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{

	}
	return r
}
