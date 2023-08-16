package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func Connect(conf Config) *gorm.DB {
	DB, err := gorm.Open(mysql.Open(conf.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic("Connection failed!")
	}
	return DB
}
