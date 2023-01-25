package config

import (
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// TODO  DB settings must checked
	d, err := gorm.Open("mysql", "name:password/tablename?charset=utf8")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
