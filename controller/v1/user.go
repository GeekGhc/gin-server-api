package v1

import (
	"gin-server-api/app"
	"gin-server-api/pkg/e"
	. "gin-server-api/service/kafka_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取用户列表
func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}

	//v,exist := c.Get(jwt.AuthUser)
	//if !exist{
	//	appG.Response(http.StatusInternalServerError, e.INVALID_PARAMS, nil)
	//	return
	//}
	//
	//user,ok := v.(*models.User)
	//if ok{
	//	fmt.Println(user)
	//}

	KafkaService.AsyncProducer("ghc")

	token := "users"
	appG.Response(http.StatusOK, e.SUCCESS, token)
}

func CreateUser(c *gin.Context) {

}
