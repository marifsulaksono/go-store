package helper

import "errors"

var (
	ErrAccDeny             = errors.New("Access Denied")
	ErrNotFound            = errors.New("Record not found")
	ErrRecDeleted          = errors.New("Record was deleted")
	ErrRecRestored         = errors.New("Record was restored")
	ErrChangeStatusProduct = errors.New("No status changed")
	ErrStockNotEnough      = errors.New("Stock product isn't enough")
	ErrUserExist           = errors.New("Username already exist")
	ErrInvalidUser         = errors.New("Unauthorized, invalid user")
	ErrDuplicateStore      = errors.New("User allowed has one store only")
)
