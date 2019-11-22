package main

import (
	"gin-server-api/models"
	"gin-server-api/pkg/logger"
	"gin-server-api/pkg/setting"
	"gin-server-api/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
	logger.Setup()
	util.Setup()
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		//nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			//"nick":    nick,
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
