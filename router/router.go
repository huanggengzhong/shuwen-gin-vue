package router

import (
	"github.com/gin-gonic/gin"
	"github.com/huanggengzhong/shuwen-gin-vue/controller"
)

func Run() {
	r := gin.Default()
	r.POST("/user/register", controller.Register)
	r.Run()
}
