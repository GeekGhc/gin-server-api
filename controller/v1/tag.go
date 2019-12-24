package v1

import (
	"github.com/gin-gonic/gin"
)

type TagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"createdBy" valid:"Required;MaxSize(100)"`
	Status    int    `form:"status" valid:"Range(0,1)"`
}

func GetTags(c *gin.Context) {
	//var (
	//	appG = app.Gin{C: c}
	//	form TagForm
	//)
	//httpCode,errCode := app.Bin
}

func AddTag(c *gin.Context) {

}
