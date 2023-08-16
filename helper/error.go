package helper

import "errors"

var (
	ErrAccDeny         = errors.New("Access Denied")
	ErrProductNotFound = errors.New("Product not found or soldout")
	ErrRecDeleted      = errors.New("Record was deleted")
	ErrRecRestored     = errors.New("Record was restored")
	ErrStockNotEnough  = errors.New("Stock product isn't enough")
	ErrUserExist       = errors.New("Username already exist")
	ErrInvalidUser     = errors.New("Unauthorized, invalid user")
	ErrInvalidSA       = errors.New("Invalid Shipping Address")
	ErrDuplicateStore  = errors.New("User allowed has one store only")
	ErrAddProductTo    = errors.New("cannot add your product's store to your cart/transaction")
)
