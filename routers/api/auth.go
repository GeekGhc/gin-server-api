package api

import (
	"gin-server-api/pkg/app"
	"gin-server-api/pkg/e"
	"gin-server-api/pkg/util"
	"gin-server-api/service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Auth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.Query("username")
	password := c.Query("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	//字段校验
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := service.Auth{Username:username,Password:password}
	isExist,err := authService.Check()
	if err != nil{
		appG.Response(http.StatusInternalServerError,e.ERROR_AUTH_CHECK_TOKEN_FAIL,nil)
		return
	}

	if !isExist{
		appG.Response(http.StatusUnauthorized,e.ERROR_AUTH_TOKEN,nil)
		return
	}

	token,err := util.GenerateToken(username,password)
	if err != nil{
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
