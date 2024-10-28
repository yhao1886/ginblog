package main

import (
	"ginblog/common"
	"ginblog/config"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	start()

	StartRouter(r)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func start() {
	config.InitConfig()

	common.InitDB()
}
