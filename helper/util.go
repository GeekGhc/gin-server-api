package helper

import (
	"gin-server-api/pkg/setting"
)

// init the util
func Setup(){
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
