package auth_service

import "gin-server-api/models"

type Auth struct {
	Username string
	Password string
}

//用户是否登录
func (a *Auth) Check() (*models.User, error) {
	return models.CheckAuth(a.Username, a.Password)
}
