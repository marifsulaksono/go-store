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
	categoryRepo := repo.NewCategoryRepository(conn)
	storeRepo := repo.NewStoreRepository(conn)
	productRepo := repo.NewProductRepository(conn)
	cartRepo := repo.NewCartRepository(conn)
	trRepo := repo.NewTransactionRepository(conn)

	userService := service.NewUserService(*userRepo)
	categoryService := service.NewCategoryService(*categoryRepo)
	storeService := service.NewStoreService(*storeRepo, *productRepo)
	productService := service.NewProductService(*productRepo)
	cartService := service.NewCartService(*cartRepo, *productRepo)
	trService := service.NewTransactionService(*trRepo, *productRepo, *userRepo)

	userController := controller.NewUserController(*userService)
	categoryController := controller.NewCategoryController(*categoryService)
	storeContoller := controller.NewStoreController(*storeService)
	productController := controller.NewProductContoller(*productService)
	cartController := controller.NewCartController(*cartService)
	trController := controller.NewTransactionController(*trService)

	// ============== Initialize Route ============

	r := mux.NewRouter()

	// ===================== Router User ======================
	r.HandleFunc("/register", userController.Register).Methods(http.MethodPost)                                                            // register new user
	r.HandleFunc("/login", userController.Login).Methods(http.MethodPost)                                                                  // login and create JSON Web Token (JWT)
	r.HandleFunc("/user/shipping_address", middleware.JWTMiddleware(userController.GetShippingAddressByUserId)).Methods(http.MethodGet)    // insert shipping address
	r.HandleFunc("/user/shipping_address", middleware.JWTMiddleware(userController.InsertShippingAddress)).Methods(http.MethodPost)        // insert shipping address
	r.HandleFunc("/user/shipping_address/{id}", middleware.JWTMiddleware(userController.UpdateShippingAddress)).Methods(http.MethodPut)    // insert shipping address
	r.HandleFunc("/user/shipping_address/{id}", middleware.JWTMiddleware(userController.DeleteShippingAddress)).Methods(http.MethodDelete) // insert shipping address

	// ==================== Router Store ====================
	r.HandleFunc("/stores", middleware.JWTMiddleware(storeContoller.GetAllStore)).Methods(http.MethodGet)
	r.HandleFunc("/store/{id}", middleware.JWTMiddleware(storeContoller.GetStoreById)).Methods(http.MethodGet)
	r.HandleFunc("/store", middleware.JWTMiddleware(storeContoller.CreateStore)).Methods(http.MethodPost)
	r.HandleFunc("/store/delete/{id}", middleware.JWTMiddleware(storeContoller.DeleteStore)).Methods(http.MethodDelete)
	r.HandleFunc("/store/{id}/delete", middleware.JWTMiddleware(storeContoller.SoftDeleteStore)).Methods(http.MethodDelete)
	r.HandleFunc("/store/{id}/restore", middleware.JWTMiddleware(storeContoller.RestoreDeletedStore)).Methods(http.MethodPut)
	r.HandleFunc("/store/{id}", middleware.JWTMiddleware(storeContoller.UpdateStore)).Methods(http.MethodPut)

	// ===================== Router Product =======================
	r.HandleFunc("/products", middleware.JWTMiddleware(productController.GetProducts)).Methods(http.MethodGet)                        // get all sales & soldout product data
	r.HandleFunc("/product/soldouts", middleware.JWTMiddleware(productController.SoldoutsProduct)).Methods(http.MethodGet)            // get all soldout product data
	r.HandleFunc("/product/search", middleware.JWTMiddleware(productController.SearchProduct)).Methods(http.MethodGet)                // get all soldout product data
	r.HandleFunc("/product/{id}/sale", middleware.JWTMiddleware(productController.ChangeProducttoSale)).Methods(http.MethodPut)       // get all sale product data
	r.HandleFunc("/product/{id}/soldout", middleware.JWTMiddleware(productController.ChangeProducttoSoldout)).Methods(http.MethodPut) // get all sale product data
	r.HandleFunc("/product/{id}", middleware.JWTMiddleware(productController.GetProductbyId)).Methods(http.MethodGet)                 // get product data by id
	r.HandleFunc("/product/delete/{id}", middleware.JWTMiddleware(productController.DeleteProduct)).Methods(http.MethodDelete)        // soft delete product data by id
	r.HandleFunc("/product/{id}/restore", middleware.JWTMiddleware(productController.RestoreDeletedProduct)).Methods(http.MethodPut)  // restore deleted product data by id
	r.HandleFunc("/product/{id}/delete", middleware.JWTMiddleware(productController.SoftDeleteProduct)).Methods(http.MethodDelete)    // hard delete product data by id
	r.HandleFunc("/product", middleware.JWTMiddleware(productController.InsertProduct)).Methods(http.MethodPost)                      // create new product data
	r.HandleFunc("/product/{id}", middleware.JWTMiddleware(productController.UpdateProduct)).Methods(http.MethodPut)                  // update product data by id

	// ==================== Router Category ====================
	r.HandleFunc("/categories", middleware.JWTMiddleware(categoryController.GetAllCategories)).Methods(http.MethodGet)

	// ==================== Router Cart ====================
	r.HandleFunc("/cart", middleware.JWTMiddleware(cartController.GetCartByUserId)).Methods(http.MethodGet)
	r.HandleFunc("/cart", middleware.JWTMiddleware(cartController.CreateCart)).Methods(http.MethodPost)
	r.HandleFunc("/cart/{id}", middleware.JWTMiddleware(cartController.UpdateCart)).Methods(http.MethodPut)
	r.HandleFunc("/cart/{id}", middleware.JWTMiddleware(cartController.DeleteCart)).Methods(http.MethodDelete)

	// ==================== Router Transaction ====================
	r.HandleFunc("/transactions", middleware.JWTMiddleware(trController.GetTransactions)).Methods(http.MethodGet)
	r.HandleFunc("/transaction/{id}", middleware.JWTMiddleware(trController.GetTransactionById)).Methods(http.MethodGet)
	r.HandleFunc("/transaction", middleware.JWTMiddleware(trController.CreateTransaction)).Methods(http.MethodPost)
	r.HandleFunc("/transaction/delete/{id}", middleware.JWTMiddleware(trController.DeleteTransaction)).Methods(http.MethodDelete)
	r.HandleFunc("/transaction/{id}/restore", middleware.JWTMiddleware(trController.RestoreDeletedTransaction)).Methods(http.MethodPut)
	r.HandleFunc("/transaction/{id}/delete", middleware.JWTMiddleware(trController.SoftDeleteTransaction)).Methods(http.MethodDelete)

	return r
}
