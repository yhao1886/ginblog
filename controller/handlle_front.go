package controller

import (
	"ginblog/common"
	"ginblog/response"
	"ginblog/service"

	"github.com/gin-gonic/gin"
)

func GetHomeInfo(ctx *gin.Context) {
	db := common.GetDB()
	result, err := service.GetFrontStatistics(db)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, result)
}
