package middleware

import (
	"errors"
	"ginblog/common"
	"ginblog/config"
	"ginblog/model"
	"ginblog/response"
	"ginblog/utils"
	"log/slog"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := common.GetDB()

		url, method := ctx.FullPath()[4:], ctx.Request.Method
		resource, err := model.GetResource(db, url, method)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				slog.Debug("[middleware-JWTAuth] resource not exist, skip jwt auth")
				ctx.Set("skip_check", true)
				ctx.Next()
				ctx.Set("skip_check", false)
				return
			}
			response.Fail(ctx, err.Error())
			return
		}

		if resource.Anonymous {
			slog.Debug("[middleware-JWTAuth] resource %s %s is anonymous, skip jwt auth", url, method)
			ctx.Set("skip_check", true)
			ctx.Next()
			ctx.Set("skip_check", false)
			return
		}

		authorization := ctx.Request.Header.Get("Authorization")
		if authorization == "" {
			response.Fail(ctx, "ErrTokenNotExist")
			return
		}

		parts := strings.Split(authorization, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Fail(ctx, "ErrTokenType")
			return
		}

		claims, err := utils.ParseToken(config.Cfg().Jwt.Secret, parts[1])
		if err != nil {
			response.Fail(ctx, "ErrTokenType")
			return
		}

		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			response.Fail(ctx, "ErrTokenRuntime")
			return
		}

		user, err := model.GetUserAuthInfoById(db, claims.UserId)
		if err != nil {
			response.Fail(ctx, "ErrTokenUserExist")
			return
		}

		// session
		session := sessions.Default(ctx)
		session.Set(utils.CTX_USER_AUTH, claims.UserId)
		session.Save()

		// gin context
		ctx.Set(utils.CTX_USER_AUTH, user)

	}
}
