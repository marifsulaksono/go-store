package entity

type Category struct {
	Id   int    `gorm:"primaryKey,autoIncrement" json:"id"`
	Name string `gorm:"not null;size:255" json:"name"`
}
