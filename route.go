package main

import (
	"gostore/middleware"
	"gostore/modules"

	"github.com/gorilla/mux"
)

func route() *mux.Router {

	r := mux.NewRouter()

	// ===================== Router Login ======================
	r.HandleFunc("/register", modules.Register).Methods("POST") // register new user
	r.HandleFunc("/login", modules.Login).Methods("POST")       // login and create JSON Web Token (JWT)

	// ===================== Router Item =======================
	r.HandleFunc("/items", middleware.JWTMiddleware(modules.GetAllItems)).Methods("GET")                // get all sales & soldout item data
	r.HandleFunc("/item/{id}", middleware.JWTMiddleware(modules.GetItembyId)).Methods("GET")            // get item data by id
	r.HandleFunc("/items", middleware.JWTMiddleware(modules.InsertItem)).Methods("POST")                // create new item data
	r.HandleFunc("/items/{id}", middleware.JWTMiddleware(modules.UpdateItem)).Methods("PUT")            // update item data by id
	r.HandleFunc("/items/sale", middleware.JWTMiddleware(modules.SalesItem)).Methods("GET")             // get all sale item data
	r.HandleFunc("/items/sale/{id}", middleware.JWTMiddleware(modules.SaleItem)).Methods("PUT")         // update item to sale
	r.HandleFunc("/items/soldout", middleware.JWTMiddleware(modules.SoldoutsItem)).Methods("GET")       // get all soldout item data
	r.HandleFunc("/items/soldout/{id}", middleware.JWTMiddleware(modules.SoldoutItem)).Methods("PUT")   // update item to soldout
	r.HandleFunc("/items/{id}", middleware.JWTMiddleware(modules.DeleteItem)).Methods("DELETE")         // delete item data by id
	r.HandleFunc("/items/category/{id}", middleware.JWTMiddleware(modules.CategoryItem)).Methods("GET") // filter items by category

	// ==================== Router Employee ====================
	r.HandleFunc("/employees", middleware.JWTMiddleware(modules.GetAllEmployees)).Methods("GET")                // get all active & inactive employee data
	r.HandleFunc("/employee/{id}", middleware.JWTMiddleware(modules.GetEmployeebyId)).Methods("GET")            // get employee data by id
	r.HandleFunc("/employees", middleware.JWTMiddleware(modules.InsertEmployee)).Methods("POST")                // create new employee data
	r.HandleFunc("/employees/{id}", middleware.JWTMiddleware(modules.UpdateEmployee)).Methods("PUT")            // update employee data by id
	r.HandleFunc("/employees/active", middleware.JWTMiddleware(modules.ActivedEmployee)).Methods("GET")         // get all active employee data
	r.HandleFunc("/employees/active/{id}", middleware.JWTMiddleware(modules.ActiveEmployee)).Methods("PUT")     // update employee to active
	r.HandleFunc("/employees/inactive", middleware.JWTMiddleware(modules.InactivedEmployee)).Methods("GET")     // get all non-active employee data
	r.HandleFunc("/employees/inactive/{id}", middleware.JWTMiddleware(modules.InactiveEmployee)).Methods("PUT") // update employee to non-active
	r.HandleFunc("/employees/{id}", middleware.JWTMiddleware(modules.DeleteEmployee)).Methods("DELETE")         // delete employee data by id

	return r
}
