package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func getDb(userName, password, host, database string, port int64) (db *gorm.DB, err error) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", userName, password, host, port, database)
	return gorm.Open("mysql", source)
}

// InitDbByInputArgs 通过参数获取mysql连接
func InitDbByInputArgs() (err error) {
	Db, err = getDb(Input.User, Input.Password, Input.Host, Input.Database, Input.Port)
	return err
}
