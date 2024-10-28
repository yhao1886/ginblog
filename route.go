package main

import (
	"ginblog/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter(r *gin.Engine) {
	base := r.Group("/api/front")
	{
		base.GET("/home", controller.GetHomeInfo)
	}
}
