package main

import (
	"gostore/controller"
	"gostore/middleware"

	"github.com/gorilla/mux"
)

func route() *mux.Router {

	r := mux.NewRouter()

	// ===================== Router Login ======================
	r.HandleFunc("/register", controller.Register).Methods("POST") // register new user
	r.HandleFunc("/login", controller.Login).Methods("POST")       // login and create JSON Web Token (JWT)

	// ===================== Router Item =======================
	r.HandleFunc("/items", middleware.JWTMiddleware(controller.GetAllItems)).Methods("GET")          // get all sales & soldout item data
	r.HandleFunc("/item/{id}", middleware.JWTMiddleware(controller.GetItembyId)).Methods("GET")      // get item data by id
	r.HandleFunc("/item", middleware.JWTMiddleware(controller.InsertItem)).Methods("POST")           // create new item data
	r.HandleFunc("/item/{id}", middleware.JWTMiddleware(controller.UpdateItem)).Methods("PUT")       // update item data by id
	r.HandleFunc("/item/sales", middleware.JWTMiddleware(controller.SalesItem)).Methods("GET")       // get all sale item data
	r.HandleFunc("/item/soldouts", middleware.JWTMiddleware(controller.SoldoutsItem)).Methods("GET") // get all soldout item data
	r.HandleFunc("/item/{id}", middleware.JWTMiddleware(controller.DeleteItem)).Methods("DELETE")    // delete item data by id

	// ==================== Router Category ====================
	r.HandleFunc("/item/category/{id}", middleware.JWTMiddleware(controller.GetAllCategory)).Methods("GET") // filter items by category

	// ==================== Router Transaction Item ====================
	r.HandleFunc("/transactionItems", middleware.JWTMiddleware(controller.GetTransactionItem)).Methods("GET")
	r.HandleFunc("/transactionItem/{id}", middleware.JWTMiddleware(controller.GetTransactionItemId)).Methods("GET")
	r.HandleFunc("/transactionItem", middleware.JWTMiddleware(controller.CreateTransactionItem)).Methods("POST")
	r.HandleFunc("/transactionItem/{id}", middleware.JWTMiddleware(controller.UpdateTransactionItem)).Methods("PUT")
	r.HandleFunc("/transactionItem/{id}", middleware.JWTMiddleware(controller.DeleteTransactionItem)).Methods("DELETE")

	// ==================== Router Transaction ====================
	r.HandleFunc("/transactions", middleware.JWTMiddleware(controller.GetTransaction)).Methods("GET")
	r.HandleFunc("/transaction/{id}", middleware.JWTMiddleware(controller.GetTransactionById)).Methods("GET")
	r.HandleFunc("/transaction", middleware.JWTMiddleware(controller.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transaction/{id}", middleware.JWTMiddleware(controller.UpdateTransaction)).Methods("PUT")
	r.HandleFunc("/transaction/{id}", middleware.JWTMiddleware(controller.DeleteTransaction)).Methods("DELETE")

	return r
}
