package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var db *gorm.DB
var mysqlInit sync.Once

func initDB() {
	_db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8&parseTime=True&loc=Local")
	db = _db
	//适用model名作为表名
	db.SingularTable(true)
	if err != nil {
		panic("连接数据库失败")
	}
}

func DBInstance() *gorm.DB {
	mysqlInit.Do(initDB)
	return db
}
