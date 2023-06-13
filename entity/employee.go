package entity

type Employee struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Role    string `json:"role"`
	Active  string `gorm:"column:isActive"`
}
