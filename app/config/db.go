package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB_user = "root"
	DB_pass = ""
	DB_name = "pretest"
	DB_dialect = "mysql"
)

var DB *gorm.DB

func Connect() *gorm.DB{
	if DB != nil{return DB}
	uri := fmt.Sprintf(fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",
		DB_user, DB_pass, DB_name))
	db,err := gorm.Open(DB_dialect, uri)
	if err != nil{panic(err.Error())}
	DB = db
	return DB
}