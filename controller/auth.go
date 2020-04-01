package controller

import (
	http2 "gin-server-api/app/http"
	"gin-server-api/helper"
	"gin-server-api/helper/e"
	"gin-server-api/service/auth_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Auth(c *gin.Context) {
	appG := http2.Gin{C: c}

	valid := validation.Validation{}

	username := c.Query("username")
	password := c.Query("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	//字段校验
	if !ok {
		http2.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	user, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	//用户是否存在
	if user == nil {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	//生成token
	token, err := helper.GenerateToken(user)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, token)
}
