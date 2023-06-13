package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	var err error
	newDB := "root:@tcp(127.0.0.1:3306)/db_store"
	DB, err = gorm.Open(mysql.Open(newDB), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection failed!")
		return nil, err
	}
	return DB, nil
}
