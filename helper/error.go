package helper

import (
	"errors"
)

const (
	InternalServerError  string = "INTERNAL_SERVER_ERROR"
	ProductNotFoundError string = "PRODUCT_NOT_FOUND"
)

var (
	ErrAccDeny        = errors.New("access Denied")
	ErrRecDeleted     = errors.New("record was deleted")
	ErrRecRestored    = errors.New("record was restored")
	ErrStockNotEnough = errors.New("stock product isn't enough")
	ErrUserExist      = errors.New("username already exist")
	ErrInvalidUser    = errors.New("unauthorized, invalid user")
	ErrInvalidSA      = errors.New("invalid Shipping Address")
	ErrDuplicateStore = errors.New("user allowed has one store only")
	ErrAddProductTo   = errors.New("cannot add your product's store to your cart/transaction")
)
