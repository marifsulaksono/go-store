package config

import (
	"fmt"
	"os"
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
