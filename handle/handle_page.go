package handle

import (
	"ginblog/common"
	"ginblog/model"
	"ginblog/response"

	"github.com/gin-gonic/gin"
)

func GetPage(ctx *gin.Context) {
	db := common.GetDB()
	pages, _, err := model.GetPageList(db)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, pages)
}
