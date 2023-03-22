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
	args := fmt.Sprintf("打印参数%s:%s:%s:%s:%s", driverName, host, port, database, username, password, charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("链接数据库失败:" + err.Error())
	}
	return db
}

func Run() {
	db := InitDB()
	defer db.Close()
}
