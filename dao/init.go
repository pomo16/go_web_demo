package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"lhx-github/go_web_demo/exceptions"
)

var db *gorm.DB

func InitDb() error {
	if database, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/goweb?charset=utf8&parseTime=True&loc=Local"); err == nil {
		database.SingularTable(true)
		db = database
		return nil
	} else {
		return exceptions.ErrDBStart
	}
}



