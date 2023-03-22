package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func Mysql(hostname string, port int, username string, password string, dbname string) {

	args := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database, charset)
	db, err := gorm.Open("mysql", args)
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}
