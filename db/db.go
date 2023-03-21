package db

import (
	"fmt"

	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "shuwen-gin-vue"
	username := "shuwen-gin-vue"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("打印参数%s:%s:%s:%s:%s")
}
