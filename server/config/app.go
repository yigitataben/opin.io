package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//dsn := "kullanici_adi:parola@tcp(localhost:3306)/veritabani_adi?charset=utf8&parseTime=True&loc=Local"
	dsn := "root:2780.Yy.5275@tcp(localhost:3306)/opinionDB?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
