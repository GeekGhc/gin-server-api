package routers

import (
	"gin-server-api/middleware/jwt"
	"gin-server-api/routers/api"
	v1 "gin-server-api/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// init the routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("auth", api.Auth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		//获取用户列表
		apiV1.GET("/users", v1.GetUsers)
		//新建用户
		apiV1.POST("/user", v1.CreateUser)
		//删除用户
		apiV1.DELETE("/users/:id", v1.DeleteUser)
	}
	return r
}
