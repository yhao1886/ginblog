package handle

import (
	"context"
	"ginblog/common"
	"ginblog/model"
	"ginblog/response"
	"ginblog/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetHomeInfo(ctx *gin.Context) {
	db := common.GetDB()
	result, err := model.GetFrontStatistics(db)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}
	response.Success(ctx, result)
}

var rctx = context.Background()

func LikeArticle(ctx *gin.Context) {
	auth, err := common.CurrentUserAuth(ctx)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	articleId, err := strconv.Atoi(ctx.Param("article_id"))
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	rdb := common.GetRdb()
	articleLikeUserKey := utils.ARTICLE_USER_LIKE_SET + strconv.Itoa(auth.ID)
	if rdb.SIsMember(rctx, articleLikeUserKey, articleId).Val() {
		rdb.SRem(rctx, articleLikeUserKey, articleId)
		rdb.HIncrBy(rctx, utils.ARTICLE_LIKE_COUNT, strconv.Itoa(articleId), -1)
	} else {
		rdb.SAdd(rctx, articleLikeUserKey, articleId)
		rdb.HIncrBy(rctx, utils.ARTICLE_LIKE_COUNT, strconv.Itoa(articleId), 1)
	}

	response.Success(ctx, nil)
}

type FAddCommentReq struct {
	ReplyUserId int    `json:"reply_user_id" form:"reply_user_id"`
	TopicId     int    `json:"topic_id" form:"topic_id"`
	Content     string `json:"content" form:"content"`
	ParentId    int    `json:"parent_id" form:"parent_id"`
	Type        int    `json:"type" form:"type" validate:"required,min=1,max=3" label:"评论类型"`
}

func SaveComment(ctx *gin.Context) {
	var req FAddCommentReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	auth, _ := common.CurrentUserAuth(ctx)
	db := common.GetDB()
	// isReview := utils.CONFIG_IS_COMMENT_REVIEW

	var comment *model.Comment
	var err error

	if req.ReplyUserId == 0 {
		comment, err = model.AddComment(db, auth.ID, req.TopicId, req.Type, req.Content, true)
	} else {
		comment, err = model.ReplyComment(db, auth.ID, req.TopicId, req.Type, req.ReplyUserId, req.ParentId, req.Content, true)
	}

	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	response.Success(ctx, comment)
}
