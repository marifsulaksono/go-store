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
	authRepo := repo.NewAuthentiocationRepo(conn)
	saRepo := repo.NewShippingAddressRepo(conn)
	categoryRepo := repo.NewCategoryRepository(conn)
	storeRepo := repo.NewStoreRepository(conn)
	productRepo := repo.NewProductRepository(conn)
	cartRepo := repo.NewCartRepository(conn)
	trRepo := repo.NewTransactionRepository(conn)
	notifRepo := repo.NewNotificationRepository(conn)

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthenticationService(authRepo, userRepo)
	saService := service.NewShippingAddressService(saRepo)
	categoryService := service.NewCategoryService(categoryRepo)
	storeService := service.NewStoreService(storeRepo, productRepo, userRepo)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(cartRepo, productRepo)
	trService := service.NewTransactionService(trRepo, saRepo, notifRepo)
	notifService := service.NewNotificationService(notifRepo)

	userController := controller.NewUserController(userService)
	authController := controller.NewAuthenticationController(authService)
	saController := controller.NewShippingAddressController(saService)
	categoryController := controller.NewCategoryController(categoryService)
	storeContoller := controller.NewStoreController(storeService)
	productController := controller.NewProductController(productService)
	cartController := controller.NewCartController(cartService)
	trController := controller.NewTransactionController(trService)
	notifController := controller.NewNotificationController(notifService)

	// ============== Initialize Route ============

	r := mux.NewRouter()

	r.Use(middleware.LoggerMiddleware)

	// ===================== Router User ======================
	r.HandleFunc("/register", userController.Register).Methods(http.MethodPost)
	r.HandleFunc("/login", authController.LoginController).Methods(http.MethodPost)
	r.HandleFunc("/logout", authController.LogoutController).Methods(http.MethodPost)
	r.HandleFunc("/refresh-token", authController.RenewAccessToken).Methods(http.MethodPost)
	r.HandleFunc("/auth/google-callback", controller.CallbackGoogleAuth).Methods(http.MethodGet)

	r.HandleFunc("/users/profile", middleware.JWTMiddleware(userController.UpdateUser)).Methods(http.MethodPut)
	r.HandleFunc("/users/password", middleware.JWTMiddleware(userController.ChangePasswordUser)).Methods(http.MethodPatch)
	r.HandleFunc("/users", middleware.JWTMiddleware(middleware.RBACAdminMiddleware(userController.DeleteUser))).Methods(http.MethodDelete)
	r.HandleFunc("/users/address", middleware.JWTMiddleware(saController.GetShippingAddressByUserId)).Methods(http.MethodGet)
	r.HandleFunc("/users/address/{id}", middleware.JWTMiddleware(saController.GetShippingAddressById)).Methods(http.MethodGet)
	r.HandleFunc("/users/address", middleware.JWTMiddleware(saController.InsertShippingAddress)).Methods(http.MethodPost)
	r.HandleFunc("/users/address/{id}", middleware.JWTMiddleware(saController.UpdateShippingAddress)).Methods(http.MethodPut)
	r.HandleFunc("/users/address/{id}", middleware.JWTMiddleware(saController.DeleteShippingAddress)).Methods(http.MethodDelete)
	r.HandleFunc("/users/{id}", userController.GetUserById).Methods(http.MethodGet)

	// ==================== Router Store ====================
	r.HandleFunc("/stores", storeContoller.GetAllStore).Methods(http.MethodGet)
	r.HandleFunc("/stores/{id}", storeContoller.GetStoreById).Methods(http.MethodGet)
	r.HandleFunc("/stores", middleware.JWTMiddleware(middleware.RBACSellerMiddleware(storeContoller.CreateStore))).Methods(http.MethodPost)
	r.HandleFunc("/stores/delete/{id}", middleware.JWTMiddleware(middleware.RBACAdminMiddleware(storeContoller.DeleteStore))).Methods(http.MethodDelete)
	r.HandleFunc("/stores/{id}/delete", middleware.JWTMiddleware(middleware.RBACSellerMiddleware(storeContoller.SoftDeleteStore))).Methods(http.MethodDelete)
	r.HandleFunc("/stores/{id}/restore", middleware.JWTMiddleware(middleware.RBACAdminMiddleware(storeContoller.RestoreDeletedStore))).Methods(http.MethodPut)
	r.HandleFunc("/stores/{id}", middleware.JWTMiddleware(middleware.RBACSellerMiddleware(storeContoller.UpdateStore))).Methods(http.MethodPut)

	// ===================== Router Product =======================
	r.HandleFunc("/products", productController.GetProducts).Methods(http.MethodGet)                                                                                  // get all sales & soldout product data
	r.HandleFunc("/products/{id}", productController.GetProductbyId).Methods(http.MethodGet)                                                                          // get product data by id
	r.HandleFunc("/products/delete/{id}", middleware.JWTMiddleware(middleware.RBACAdminMiddleware(productController.DeleteProduct))).Methods(http.MethodDelete)       // soft delete product data by id
	r.HandleFunc("/products/{id}/restore", middleware.JWTMiddleware(middleware.RBACAdminMiddleware(productController.RestoreDeletedProduct))).Methods(http.MethodPut) // restore deleted product data by id
	r.HandleFunc("/products/{id}/delete", middleware.JWTMiddleware(middleware.RBACSellerMiddleware(productController.SoftDeleteProduct))).Methods(http.MethodDelete)  // hard delete product data by id
	r.HandleFunc("/products", middleware.JWTMiddleware(middleware.RBACSellerMiddleware(productController.InsertProduct))).Methods(http.MethodPost)                    // create new product data
	r.HandleFunc("/products/{id}", middleware.JWTMiddleware(middleware.RBACSellerMiddleware(productController.UpdateProduct))).Methods(http.MethodPut)                // update product data by id

	// ==================== Router Category ====================
	r.HandleFunc("/categories", categoryController.GetAllCategories).Methods(http.MethodGet)
	r.HandleFunc("/categories/{id}", categoryController.GetCategoryById).Methods(http.MethodGet)
	r.HandleFunc("/categories", middleware.JWTMiddleware(middleware.RBACAdminMiddleware(categoryController.InsertCategory))).Methods(http.MethodPost)
	r.HandleFunc("/categories/{id}", middleware.JWTMiddleware(middleware.RBACAdminMiddleware(categoryController.UpdateCategory))).Methods(http.MethodPut)
	r.HandleFunc("/categories/{id}", middleware.JWTMiddleware(middleware.RBACAdminMiddleware(categoryController.DeleteCategory))).Methods(http.MethodDelete)

	// ==================== Router Cart ====================
	r.HandleFunc("/carts", middleware.JWTMiddleware(cartController.GetCartByUserId)).Methods(http.MethodGet)
	r.HandleFunc("/carts", middleware.JWTMiddleware(cartController.CreateCart)).Methods(http.MethodPost)
	r.HandleFunc("/carts/{id}", middleware.JWTMiddleware(cartController.UpdateCart)).Methods(http.MethodPut)
	r.HandleFunc("/carts/{id}", middleware.JWTMiddleware(cartController.DeleteCart)).Methods(http.MethodDelete)

	// ==================== Router Transaction ====================
	r.HandleFunc("/transactions", middleware.JWTMiddleware(trController.GetTransactions)).Methods(http.MethodGet)
	r.HandleFunc("/transactions/{id}", middleware.JWTMiddleware(trController.GetTransactionById)).Methods(http.MethodGet)
	r.HandleFunc("/transactions", middleware.JWTMiddleware(trController.CreateTransaction)).Methods(http.MethodPost)

	// ==================== Router Notification ====================
	r.HandleFunc("/notifications", middleware.JWTMiddleware(notifController.GetAllNotifications)).Methods(http.MethodGet)

	return r
}
