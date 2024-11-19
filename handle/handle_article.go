package handle

import (
	"ginblog/common"
	"ginblog/model"
	"ginblog/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PageQuery struct {
	Page    int    `form:"page_num"`
	Size    int    `form:"page_size"`
	Keyword string `form:"keyword"`
}

type FArticleQuery struct {
	PageQuery
	CategoryId int `form:"category_id"`
	TagId      int `form:"tag_id"`
}

func GetArticles(ctx *gin.Context) {
	db := common.GetDB()
	var params FArticleQuery
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	list, _, err := model.GetArticleList(db, params.Page, params.Size, params.CategoryId, params.TagId)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	response.Success(ctx, list)
}

func GetArticleInfo(ctx *gin.Context) {
	db := common.GetDB()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	data, err := model.GetBlogArticle(db, id)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, data)
}
