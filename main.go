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
	r.HandleFunc("/items", modules.GetAllItems).Methods("GET")                // get all sales & soldout item data
	r.HandleFunc("/item/{id}", modules.GetItembyId).Methods("GET")            // get item data by id
	r.HandleFunc("/items", modules.InsertItem).Methods("POST")                // create new item data
	r.HandleFunc("/items/{id}", modules.UpdateItem).Methods("PUT")            // update item data by id
	r.HandleFunc("/items/sale", modules.SalesItem).Methods("GET")             // get all sale item data
	r.HandleFunc("/items/sale/{id}", modules.SaleItem).Methods("PUT")         // update item to sale
	r.HandleFunc("/items/soldout", modules.SoldoutsItem).Methods("GET")       // get all soldout item data
	r.HandleFunc("/items/soldout/{id}", modules.SoldoutItem).Methods("PUT")   // update item to soldout
	r.HandleFunc("/items/{id}", modules.DeleteItem).Methods("DELETE")         // delete item data by id
	r.HandleFunc("/items/category/{id}", modules.CategoryItem).Methods("GET") // filter items by category

	// ==================== Handler Employee ====================
	r.HandleFunc("/employees", modules.GetAllEmployees).Methods("GET")                // get all active & inactive employee data
	r.HandleFunc("/employee/{id}", modules.GetEmployeebyId).Methods("GET")            // get employee data by id
	r.HandleFunc("/employees", modules.InsertEmployee).Methods("POST")                // create new employee data
	r.HandleFunc("/employees/{id}", modules.UpdateEmployee).Methods("PUT")            // update employee data by id
	r.HandleFunc("/employees/active", modules.ActivedEmployee).Methods("GET")         // get all active employee data
	r.HandleFunc("/employees/active/{id}", modules.ActiveEmployee).Methods("PUT")     // update employee to active
	r.HandleFunc("/employees/inactive", modules.InactivedEmployee).Methods("GET")     // get all non-active employee data
	r.HandleFunc("/employees/inactive/{id}", modules.InactiveEmployee).Methods("PUT") // update employee to non-active
	r.HandleFunc("/employees/{id}", modules.DeleteEmployee).Methods("DELETE")         // delete employee data by id

	// ==================== Start Server ====================
	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", r)
}
