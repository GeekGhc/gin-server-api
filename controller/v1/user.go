package v1

import (
	"fmt"
	"gin-server-api/middleware/jwt"
	"gin-server-api/models"
	"gin-server-api/pkg/app"
	"gin-server-api/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取用户列表
func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}

	v,exist := c.Get(jwt.AuthUser)
	if !exist{
		appG.Response(http.StatusInternalServerError, e.INVALID_PARAMS, nil)
		return
	}

	user,ok := v.(*models.User)
	if ok{
		fmt.Println(user)
	}

	token := "users"
	appG.Response(http.StatusOK, e.SUCCESS, token)
}

func CreateUser(c *gin.Context) {

}
