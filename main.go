package main

import (
	"github.com/gin-gonic/gin"
	"github.com/huanggengzhong/shuwen-gin-vue/controller"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.GET("/api/user/rigister", controller.Register)
	r.Run()
}
