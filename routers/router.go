package routers

import (
	"gin-server-api/controller"
	v1 "gin-server-api/controller/v1"
	"gin-server-api/middleware"
	"github.com/gin-gonic/gin"
)

// init the routing information
func InitRouter(r *gin.Engine) {
	r.GET("auth", controller.Auth)
	//心跳url
	r.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		//获取用户列表
		apiV1.GET("/users", v1.GetUsers)
		//新建用户
		apiV1.POST("/user", v1.CreateUser)

		//标签列表
		apiV1.GET("/tags", v1.GetTags)
		//新建标签
		apiV1.POST("/tag", v1.AddTag)
	}
}
