package helper

import (
	"gin-server-api/app/setting"
)

// init the util
func Setup(){
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
