package main

import (
	"gostore/config"

	"fmt"
	"net/http"
)

func main() {
	conn := config.Connect()

	// ==================== Start Server ====================
	fmt.Println("Server started at localhost:49999")
	http.ListenAndServe(":49999", routeInit(conn))
}
