package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huanggengzhong/shuwen-gin-vue/util"
	"github.com/shuwenhe/utils"
)

func Register(ctx *gin.Context) {
	// 1.获取参数
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")
	phone := ctx.PostForm("phone")
	// 2.验证
	if len(phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号长度不对",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能少于6位",
		})
		return
	}
	//没有名字自动填充一个
	if len(name) == 0 {
		name := util.RandomString(10)
		fmt.Println("自己添加name=", name)
	}
	// 3.返回结果

	ctx.JSON(utils.NewSucc("注册成功", gin.H{
		"msg": "success",
	}))
}
