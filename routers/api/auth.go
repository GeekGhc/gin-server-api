package api

import (
	"gin-server-api/pkg/app"
	"gin-server-api/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": "xxxx",
	})
}
