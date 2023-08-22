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
	// ============== Dependency Injection ============
	userRepo := repo.NewUserRepository(conn)
	categoryRepo := repo.NewCategoryRepository(conn)
	storeRepo := repo.NewStoreRepository(conn)
	productRepo := repo.NewProductRepository(conn)
	cartRepo := repo.NewCartRepository(conn)
	trRepo := repo.NewTransactionRepository(conn)

	userService := service.NewUserService(userRepo)
	categoryService := service.NewCategoryService(*categoryRepo)
	storeService := service.NewStoreService(storeRepo, productRepo)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(cartRepo, productRepo)
	trService := service.NewTransactionService(trRepo, productRepo, userRepo)

	userController := controller.NewUserController(userService)
	categoryController := controller.NewCategoryController(*categoryService)
	storeContoller := controller.NewStoreController(storeService)
	productController := controller.NewProductController(productService)
	cartController := controller.NewCartController(cartService)
	trController := controller.NewTransactionController(trService)

	// ============== Initialize Route ============

	r := mux.NewRouter()

	// ===================== Router User ======================
	r.HandleFunc("/register", userController.Register).Methods(http.MethodPost)
	r.HandleFunc("/login", userController.Login).Methods(http.MethodPost)
	r.HandleFunc("/user/profile", middleware.JWTMiddleware(userController.UpdateUser)).Methods(http.MethodPut)
	r.HandleFunc("/user/password", middleware.JWTMiddleware(userController.ChangePasswordUser)).Methods(http.MethodPatch)
	r.HandleFunc("/user/address", middleware.JWTMiddleware(userController.GetShippingAddressByUserId)).Methods(http.MethodGet)
	r.HandleFunc("/user/address", middleware.JWTMiddleware(userController.InsertShippingAddress)).Methods(http.MethodPost)
	r.HandleFunc("/user/address/{id}", middleware.JWTMiddleware(userController.UpdateShippingAddress)).Methods(http.MethodPut)
	r.HandleFunc("/user/address/{id}", middleware.JWTMiddleware(userController.DeleteShippingAddress)).Methods(http.MethodDelete)
	r.HandleFunc("/user", middleware.JWTMiddleware(userController.DeleteUser)).Methods(http.MethodDelete)

	// ==================== Router Store ====================
	r.HandleFunc("/stores", storeContoller.GetAllStore).Methods(http.MethodGet)
	r.HandleFunc("/stores/{id}", storeContoller.GetStoreById).Methods(http.MethodGet)
	r.HandleFunc("/stores", middleware.JWTMiddleware(storeContoller.CreateStore)).Methods(http.MethodPost)
	r.HandleFunc("/stores/delete/{id}", middleware.JWTMiddleware(storeContoller.DeleteStore)).Methods(http.MethodDelete)
	r.HandleFunc("/stores/{id}/delete", middleware.JWTMiddleware(storeContoller.SoftDeleteStore)).Methods(http.MethodDelete)
	r.HandleFunc("/stores/{id}/restore", middleware.JWTMiddleware(storeContoller.RestoreDeletedStore)).Methods(http.MethodPut)
	r.HandleFunc("/stores/{id}", middleware.JWTMiddleware(storeContoller.UpdateStore)).Methods(http.MethodPut)

	// ===================== Router Product =======================
	r.HandleFunc("/products/search", productController.GetProducts).Methods(http.MethodGet)                                           // get all sales & soldout product data
	r.HandleFunc("/products/{id}", productController.GetProductbyId).Methods(http.MethodGet)                                          // get product data by id
	r.HandleFunc("/products/delete/{id}", middleware.JWTMiddleware(productController.DeleteProduct)).Methods(http.MethodDelete)       // soft delete product data by id
	r.HandleFunc("/products/{id}/restore", middleware.JWTMiddleware(productController.RestoreDeletedProduct)).Methods(http.MethodPut) // restore deleted product data by id
	r.HandleFunc("/products/{id}/delete", middleware.JWTMiddleware(productController.SoftDeleteProduct)).Methods(http.MethodDelete)   // hard delete product data by id
	r.HandleFunc("/products", middleware.JWTMiddleware(productController.InsertProduct)).Methods(http.MethodPost)                     // create new product data
	r.HandleFunc("/products/{id}", middleware.JWTMiddleware(productController.UpdateProduct)).Methods(http.MethodPut)                 // update product data by id

	// ==================== Router Category ====================
	r.HandleFunc("/categories", middleware.JWTMiddleware(categoryController.GetAllCategories)).Methods(http.MethodGet)

	// ==================== Router Cart ====================
	r.HandleFunc("/carts", middleware.JWTMiddleware(cartController.GetCartByUserId)).Methods(http.MethodGet)
	r.HandleFunc("/carts", middleware.JWTMiddleware(cartController.CreateCart)).Methods(http.MethodPost)
	r.HandleFunc("/carts/{id}", middleware.JWTMiddleware(cartController.UpdateCart)).Methods(http.MethodPut)
	r.HandleFunc("/carts/{id}", middleware.JWTMiddleware(cartController.DeleteCart)).Methods(http.MethodDelete)

	// ==================== Router Transaction ====================
	r.HandleFunc("/transactions", middleware.JWTMiddleware(trController.GetTransactions)).Methods(http.MethodGet)
	r.HandleFunc("/transactions/{id}", middleware.JWTMiddleware(trController.GetTransactionById)).Methods(http.MethodGet)
	r.HandleFunc("/transactions", middleware.JWTMiddleware(trController.CreateTransaction)).Methods(http.MethodPost)

	return r
}
