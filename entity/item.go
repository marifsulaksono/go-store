package entity

type Item struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Stock      int      `json:"stock"`
	Price      int      `json:"price"`
	Sale       int      `gorm:"column:isSale" json:"isSale"`
	CategoryId int      `json:"categoryId"`
	Category   Category `json:"category"`
	// Category   Category `gorm:"foreignKey:IdCategory" json:"category"` // inisialisasi foreignkey pada gorm
}

type ItemResponse struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Stock      int      `json:"stock"`
	Price      int      `json:"price"`
	Sale       int      `gorm:"column:isSale" json:"isSale"`
	CategoryId int      `json:"-"`
	Category   Category `json:"category"`
}

func (ItemResponse) TableName() string {
	return "items"
}
