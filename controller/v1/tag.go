package v1

import (
	"gin-server-api/app"
	"gin-server-api/pkg/e"
	"gin-server-api/service/tag_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TagForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	CreatedBy int    `form:"createdBy" valid:"Required"`
	Status    int    `form:"status" valid:"Range(0,1)"`
}

func GetTags(c *gin.Context) {

}

//添加标签
func AddTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form TagForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := tag_service.Tag{
		Name:      form.Name,
		CreatedBy: form.CreatedBy,
		Status:    form.Status,
	}
	exist, err := tagService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	//已经存在该标签
	if exist {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}

	err = tagService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
