package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, httpStatus int, code int, data any, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code":    code,
		"data":    data,
		"message": msg,
	})
}

func Success(ctx *gin.Context, data any) {
	Response(ctx, http.StatusOK, 0, data, "OK")
}

func Fail(ctx *gin.Context, msg string) {
	Response(ctx, http.StatusOK, 400, nil, msg)
}
