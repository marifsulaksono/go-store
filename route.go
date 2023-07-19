package main

import (
	"gostore/controller"
	"gostore/middleware"
	"gostore/repo"
	"gostore/service"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func routeInit(conn *gorm.DB) *mux.Router {
	// ============== Initialize Dependency ============
	userRepo := repo.NewUserRepository(conn)
	userService := service.NewUserService(*userRepo)
	userController := controller.NewUserController(*userService)

	categoryRepo := repo.NewCategoryRepository(conn)
	categoryService := service.NewCategoryService(*categoryRepo)
	categoryController := controller.NewCategoryController(*categoryService)

	itemRepo := repo.NewItemRepository(conn)
	itemService := service.NewItemService(*itemRepo)
	itemController := controller.NewItemContoller(*itemService)

	trItemRepo := repo.NewTransactionItemRepository(conn)
	trItemService := service.NewTransactionItemService(*trItemRepo)
	trItemController := controller.NewTransactionItemController(*trItemService)

	trRepo := repo.NewTransactionRepository(conn)
	trService := service.NewTransactionService(*trRepo)
	trController := controller.NewTransactionController(*trService)

	// ============== Initialize Route ============

	r := mux.NewRouter()

	// ===================== Router Login ======================
	r.HandleFunc("/register", userController.Register).Methods(http.MethodPost) // register new user
	r.HandleFunc("/login", userController.Login).Methods(http.MethodPost)       // login and create JSON Web Token (JWT)

	// ===================== Router Item =======================
	r.HandleFunc("/items", middleware.JWTMiddleware(itemController.GetItems)).Methods(http.MethodGet)                       // get all sales & soldout item data
	r.HandleFunc("/item/sales", middleware.JWTMiddleware(itemController.SalesItem)).Methods(http.MethodGet)                 // get all sale item data
	r.HandleFunc("/item/soldouts", middleware.JWTMiddleware(itemController.SoldoutsItem)).Methods(http.MethodGet)           // get all soldout item data
	r.HandleFunc("/item/{id}", middleware.JWTMiddleware(itemController.GetItembyId)).Methods(http.MethodGet)                // get item data by id
	r.HandleFunc("/item/restore/{id}", middleware.JWTMiddleware(itemController.RestoreDeletedItem)).Methods(http.MethodPut) // restore deleted item data by id
	r.HandleFunc("/item/delete/{id}", middleware.JWTMiddleware(itemController.DeleteItem)).Methods(http.MethodDelete)       // hard delete item data by id
	r.HandleFunc("/item", middleware.JWTMiddleware(itemController.InsertItem)).Methods(http.MethodPost)                     // create new item data
	r.HandleFunc("/item/{id}", middleware.JWTMiddleware(itemController.UpdateItem)).Methods(http.MethodPut)                 // update item data by id
	r.HandleFunc("/item/{id}", middleware.JWTMiddleware(itemController.SoftDeleteItem)).Methods(http.MethodDelete)          // soft delete item data by id

	// ==================== Router Category ====================
	r.HandleFunc("/categories", middleware.JWTMiddleware(categoryController.GetAllCategories)).Methods(http.MethodGet) // filter items by category

	// ==================== Router Transaction Item ====================
	r.HandleFunc("/transactionItems", middleware.JWTMiddleware(trItemController.GetTransactionItems)).Methods(http.MethodGet)
	r.HandleFunc("/transactionItem/{id}", middleware.JWTMiddleware(trItemController.GetTransactionItemId)).Methods(http.MethodGet)
	r.HandleFunc("/transactionItem", middleware.JWTMiddleware(trItemController.CreateTransactionItem)).Methods(http.MethodPost)
	r.HandleFunc("/transactionItem/restore/{id}", middleware.JWTMiddleware(trItemController.RestoreDeletedTransactionItem)).Methods(http.MethodPut)
	r.HandleFunc("/transactionItem/{id}", middleware.JWTMiddleware(trItemController.UpdateTransactionItem)).Methods(http.MethodPut)
	r.HandleFunc("/transactionItem/delete/{id}", middleware.JWTMiddleware(trItemController.DeleteTransactionItem)).Methods(http.MethodDelete)
	r.HandleFunc("/transactionItem/{id}", middleware.JWTMiddleware(trItemController.SoftDeleteTransactionItem)).Methods(http.MethodDelete)

	// ==================== Router Transaction ====================
	r.HandleFunc("/transactions", middleware.JWTMiddleware(trController.GetTransactions)).Methods(http.MethodGet)
	r.HandleFunc("/transaction/{id}", middleware.JWTMiddleware(trController.GetTransactionById)).Methods(http.MethodGet)
	r.HandleFunc("/transaction", middleware.JWTMiddleware(trController.CreateTransaction)).Methods(http.MethodPost)
	r.HandleFunc("/transaction/restore/{id}", middleware.JWTMiddleware(trController.RestoreDeletedTransaction)).Methods(http.MethodPut)
	r.HandleFunc("/transaction/{id}", middleware.JWTMiddleware(trController.UpdateTransaction)).Methods(http.MethodPut)
	r.HandleFunc("/transaction/delete/{id}", middleware.JWTMiddleware(trController.DeleteTransaction)).Methods(http.MethodDelete)
	r.HandleFunc("/transaction/{id}", middleware.JWTMiddleware(trController.SoftDeleteTransaction)).Methods(http.MethodDelete)

	return r
}
