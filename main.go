package main

import (
	"gostore/config"
	"gostore/modules"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// ==================== Fungsi main ====================

func main() {
	config.Connect()
	r := mux.NewRouter()

	// ===================== Handler Login ======================
	r.HandleFunc("/register", modules.Register).Methods("POST") // register new user
	r.HandleFunc("/login", modules.Login).Methods("POST")       // login and create JSON Web Token (JWT)

	// ===================== Handler Item =======================
	r.HandleFunc("/items", modules.JWTMiddleware(modules.GetAllItems)).Methods("GET")                // get all sales & soldout item data
	r.HandleFunc("/item/{id}", modules.JWTMiddleware(modules.GetItembyId)).Methods("GET")            // get item data by id
	r.HandleFunc("/items", modules.JWTMiddleware(modules.InsertItem)).Methods("POST")                // create new item data
	r.HandleFunc("/items/{id}", modules.JWTMiddleware(modules.UpdateItem)).Methods("PUT")            // update item data by id
	r.HandleFunc("/items/sale", modules.JWTMiddleware(modules.SalesItem)).Methods("GET")             // get all sale item data
	r.HandleFunc("/items/sale/{id}", modules.JWTMiddleware(modules.SaleItem)).Methods("PUT")         // update item to sale
	r.HandleFunc("/items/soldout", modules.JWTMiddleware(modules.SoldoutsItem)).Methods("GET")       // get all soldout item data
	r.HandleFunc("/items/soldout/{id}", modules.JWTMiddleware(modules.SoldoutItem)).Methods("PUT")   // update item to soldout
	r.HandleFunc("/items/{id}", modules.JWTMiddleware(modules.DeleteItem)).Methods("DELETE")         // delete item data by id
	r.HandleFunc("/items/category/{id}", modules.JWTMiddleware(modules.CategoryItem)).Methods("GET") // filter items by category

	// ==================== Handler Employee ====================
	r.HandleFunc("/employees", modules.JWTMiddleware(modules.GetAllEmployees)).Methods("GET")                // get all active & inactive employee data
	r.HandleFunc("/employee/{id}", modules.JWTMiddleware(modules.GetEmployeebyId)).Methods("GET")            // get employee data by id
	r.HandleFunc("/employees", modules.JWTMiddleware(modules.InsertEmployee)).Methods("POST")                // create new employee data
	r.HandleFunc("/employees/{id}", modules.JWTMiddleware(modules.UpdateEmployee)).Methods("PUT")            // update employee data by id
	r.HandleFunc("/employees/active", modules.JWTMiddleware(modules.ActivedEmployee)).Methods("GET")         // get all active employee data
	r.HandleFunc("/employees/active/{id}", modules.JWTMiddleware(modules.ActiveEmployee)).Methods("PUT")     // update employee to active
	r.HandleFunc("/employees/inactive", modules.JWTMiddleware(modules.InactivedEmployee)).Methods("GET")     // get all non-active employee data
	r.HandleFunc("/employees/inactive/{id}", modules.JWTMiddleware(modules.InactiveEmployee)).Methods("PUT") // update employee to non-active
	r.HandleFunc("/employees/{id}", modules.JWTMiddleware(modules.DeleteEmployee)).Methods("DELETE")         // delete employee data by id

	// ==================== Start Server ====================
	fmt.Println("Server started at localhost:49999")
	http.ListenAndServe(":49999", r)
}
