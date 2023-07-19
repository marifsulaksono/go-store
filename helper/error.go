package helper

import "errors"

var (
	ErrAccDeny          = errors.New("Access Denied!")
	ErrNotFound         = errors.New("Record not found")
	ErrRecDeleted       = errors.New("Record was deleted")
	ErrRecRestored      = errors.New("Record was restored")
	ErrChangeStatusItem = errors.New("Already change")
	ErrUserExist        = errors.New("Username already exist")
)
