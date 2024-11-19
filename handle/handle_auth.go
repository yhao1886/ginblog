package handle

import (
	"ginblog/common"
	"ginblog/config"
	"ginblog/model"
	"ginblog/response"
	"ginblog/utils"
	"log/slog"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginVO struct {
	model.UserInfo

	// 点赞 Set: 用于记录用户点赞过的文章, 评论
	ArticleLikeSet []string `json:"article_like_set"`
	CommentLikeSet []string `json:"comment_like_set"`
	Token          string   `json:"token"`
}

func Login(ctx *gin.Context) {
	var req loginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	db := common.GetDB()
	userAuth, err := model.GetUserAuthInfoByName(db, req.Username)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	if !utils.BcryptCheck(req.Password, userAuth.Password) {
		response.Fail(ctx, "password error")
		return
	}

	userInfo, err := model.GetUserInfoById(db, userAuth.UserInfoId)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	ipAddress := utils.GetIpAddress(ctx)
	ipSource := utils.IP.GetIpSourceSimpleIdle(ipAddress)

	cfg := config.Cfg()

	token, err := utils.GetToken(cfg.Jwt.Secret, cfg.Issuer, req.Username, int(cfg.Expire), userAuth.ID)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	if err = model.UpdateUserLoginInfo(db, userAuth.ID, ipAddress, ipSource); err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	slog.Info("用户登录成功: " + userAuth.Username)

	session := sessions.Default(ctx)
	session.Set(utils.CTX_USER_AUTH, userAuth.ID)
	session.Save()

	response.Success(ctx, LoginVO{
		UserInfo: userInfo,
		Token:    token,
	})

}

func GetInfo(ctx *gin.Context) {
	user, err := common.CurrentUserAuth(ctx)
	if err != nil {
		response.Fail(ctx, err.Error())
		return
	}

	userInfoVO := model.UserInfoVO{
		UserInfo:       *user.UserInfo,
		ArticleLikeSet: make([]string, 0),
		CommentLikeSet: make([]string, 0),
	}
	response.Success(ctx, userInfoVO)

}
