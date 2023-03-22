package main

import (
	"github.com/huanggengzhong/shuwen-gin-vue/db"
	"github.com/huanggengzhong/shuwen-gin-vue/router"
)

func main() {
	db.Run()
	router.Run()
}
