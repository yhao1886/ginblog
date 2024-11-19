package main

import (
	"ginblog/handle"
	"ginblog/middleware"

	"github.com/gin-gonic/gin"
)

func StartBaseRouter(r *gin.Engine) {
	base := r.Group("/api")
	{
		base.POST("/login", handle.Login)
	}
}

func StartRouter(r *gin.Engine) {
	base := r.Group("/api/front")
	{
		base.GET("/home", handle.GetHomeInfo)
		base.GET("/page", handle.GetPage)
	}

	article := base.Group("/article")
	{
		article.GET("/list", handle.GetArticles)
		article.GET("/:id", handle.GetArticleInfo)
	}

	category := base.Group("/category")
	{
		category.GET("/list", handle.GetCategorys)
	}

	tag := base.Group("/tag")
	{
		tag.GET("list", handle.GetTags)
	}

	comment := base.Group("/comment")
	{
		comment.GET("/list")
	}

	base.Use(middleware.JWTAuth())
	{
		base.GET("/user/info", handle.GetInfo)
		base.GET("/article/like/:article_id", handle.LikeArticle)
		base.POST("/comment", handle.SaveComment)
	}
}
