package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/utils"
)

func Register(ctx *gin.Context) {
	ctx.JSON(utils.NewSucc())
}
