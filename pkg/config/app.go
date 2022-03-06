package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//d, err := gorm.Open("mysql", "appuser:appuser@tcp(host.docker.internal:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", "appuser:appuser@tcp(db:3300)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
