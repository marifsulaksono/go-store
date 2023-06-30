package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	newDB := "root:@tcp(127.0.0.1:3306)/db_store?parseTime=true"
	DB, err = gorm.Open(mysql.Open(newDB), &gorm.Config{})
	if err != nil {
		panic("Connection failed!")
	}
}
