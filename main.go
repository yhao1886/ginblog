package main

import (
	"ginblog/common"
	"ginblog/config"
	"ginblog/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	start()

	conf := config.Cfg()

	r.Use(middleware.WithCookieStore(conf.Session.Name, conf.Session.Secret, conf.Age))

	StartBaseRouter(r)
	StartRouter(r)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func start() {
	config.InitConfig()

	common.InitDB()

	common.InitRedis()
}
