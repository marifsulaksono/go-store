package main

import (
	"gostore/config"
	"log"

	"fmt"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env files : %v", err)
	}

	conf := config.GetConfig()
	conn := config.Connect(conf)
	// config.AutoMigrate(conn)

	// ==================== Start Server ====================
	fmt.Printf("Server starting at localhost:%v ...\n", conf.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", conf.Port), routeInit(conn))
	if err != nil {
		log.Fatalf("Error starting server : %v", err)
	}
}
