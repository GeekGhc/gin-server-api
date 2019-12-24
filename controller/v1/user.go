package v1

import (
	"gin-server-api/pkg/app"
	"gin-server-api/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取用户列表
func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}

	token := "users"
	appG.Response(http.StatusOK, e.SUCCESS, token)
}

func CreateUser(c *gin.Context) {

}
