package entity

type Cart struct {
	Id        int                        `gorm:"primaryKey,autoIncrement" json:"id"`
	UserId    int                        `gorm:"not null" json:"-"`
	ProductId *int                       `gorm:"not null" json:"product_id"`
	Product   ProductTransactionResponse `gorm:"-:migration" json:"product"`
	Qty       *int                       `gorm:"not null" json:"qty"`
}

// {
// 	"product_id": 1,
// 	"qty": 10
// }
