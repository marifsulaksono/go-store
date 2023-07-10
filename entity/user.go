package entity

import "gostore/config"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func GetUserValid(username string) (User, error) {
	var userLogin User
	err := config.DB.Where("username = ?", username).First(&userLogin).Error
	return userLogin, err
}
