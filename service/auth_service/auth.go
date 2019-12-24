package auth_service

import "gin-server-api/models"

type Auth struct {
	Username string
	Password string
}

//用户是否登录
func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}
