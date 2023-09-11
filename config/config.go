package config

import (
	"fmt"
	"gostore/entity"
	"log"
	"os"

	"gorm.io/gorm"
)

const (
	dbUsername = "DB_USERNAME"
	dbPassword = "DB_PASSWORD"
	dbHost     = "DB_HOST"
	dbName     = "DB_NAME"
	serverPort = "SERVER_PORT"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func GetConfig() Config {
	return Config{
		Port: os.Getenv(serverPort),
		DatabaseURL: fmt.Sprintf("%v:%v@tcp(%v)/%v", os.Getenv(dbUsername),
			os.Getenv(dbPassword), os.Getenv(dbHost), os.Getenv(dbName)),
	}
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.ShippingAddress{},
		&entity.Product{},
		&entity.Category{},
		&entity.Store{},
		&entity.Cart{},
		&entity.Transaction{},
		&entity.TransactionItem{},
	)

	if err != nil {
		log.Fatalf("Migration Failed. Error : %v", err)
	}

	log.Println("Migration Success....")
}
