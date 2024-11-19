package common

import (
	"errors"
	"ginblog/model"
	"ginblog/utils"
	"log/slog"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CurrentUserAuth(ctx *gin.Context) (*model.UserAuth, error) {
	key := utils.CTX_USER_AUTH

	if cache, exist := ctx.Get(key); exist && cache != nil {
		slog.Debug("[Func-CurrentUserAuth] get from cache: " + cache.(*model.UserAuth).Username)
	}

	session := sessions.Default(ctx)
	id := session.Get(key)
	if id == nil {
		return nil, errors.New("session 中没有 user_auth_id")
	}

	db := GetDB()
	user, err := model.GetUserAuthInfoById(db, id.(int))
	if err != nil {
		return nil, err
	}

	ctx.Set(key, user)

	return &user, nil
}
