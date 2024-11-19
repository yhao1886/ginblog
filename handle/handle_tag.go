package handle

import (
	"ginblog/common"
	"ginblog/model"
	"ginblog/response"

	"github.com/gin-gonic/gin"
)

func GetTags(ctx *gin.Context) {
	db := common.GetDB()
	result, err := model.GetTaglist(db)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, result)
}
