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

	// ===================== Handler Item ======================
	r.HandleFunc("/items", modules.Middleware(modules.GetAllItems)).Methods("GET")                // get all sales & soldout item data
	r.HandleFunc("/item/{id}", modules.Middleware(modules.GetItembyId)).Methods("GET")            // get item data by id
	r.HandleFunc("/items", modules.Middleware(modules.InsertItem)).Methods("POST")                // create new item data
	r.HandleFunc("/items/{id}", modules.Middleware(modules.UpdateItem)).Methods("PUT")            // update item data by id
	r.HandleFunc("/items/sale", modules.Middleware(modules.SalesItem)).Methods("GET")             // get all sale item data
	r.HandleFunc("/items/sale/{id}", modules.Middleware(modules.SaleItem)).Methods("PUT")         // update item to sale
	r.HandleFunc("/items/soldout", modules.Middleware(modules.SoldoutsItem)).Methods("GET")       // get all soldout item data
	r.HandleFunc("/items/soldout/{id}", modules.Middleware(modules.SoldoutItem)).Methods("PUT")   // update item to soldout
	r.HandleFunc("/items/{id}", modules.Middleware(modules.DeleteItem)).Methods("DELETE")         // delete item data by id
	r.HandleFunc("/items/category/{id}", modules.Middleware(modules.CategoryItem)).Methods("GET") // filter items by category

	// ==================== Handler Employee ====================
	r.HandleFunc("/employees", modules.Middleware(modules.GetAllEmployees)).Methods("GET")                // get all active & inactive employee data
	r.HandleFunc("/employee/{id}", modules.Middleware(modules.GetEmployeebyId)).Methods("GET")            // get employee data by id
	r.HandleFunc("/employees", modules.Middleware(modules.InsertEmployee)).Methods("POST")                // create new employee data
	r.HandleFunc("/employees/{id}", modules.Middleware(modules.UpdateEmployee)).Methods("PUT")            // update employee data by id
	r.HandleFunc("/employees/active", modules.Middleware(modules.ActivedEmployee)).Methods("GET")         // get all active employee data
	r.HandleFunc("/employees/active/{id}", modules.Middleware(modules.ActiveEmployee)).Methods("PUT")     // update employee to active
	r.HandleFunc("/employees/inactive", modules.Middleware(modules.InactivedEmployee)).Methods("GET")     // get all non-active employee data
	r.HandleFunc("/employees/inactive/{id}", modules.Middleware(modules.InactiveEmployee)).Methods("PUT") // update employee to non-active
	r.HandleFunc("/employees/{id}", modules.Middleware(modules.DeleteEmployee)).Methods("DELETE")         // delete employee data by id

	// ==================== Start Server ====================
	fmt.Println("Server started at localhost:49999")
	http.ListenAndServe(":49999", r)
}
