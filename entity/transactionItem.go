package entity

type TransactionItem struct {
	Id            int `json:"id"`
	TransactionId int `json:"transaction_id"`
	ItemId        int `json:"item_id"`
	Qty           int `json:"qty"`
	Price         int `json:"price"`
	Subtotal      int `json:"subtotal"`
}

type TransactionItemResponse struct {
	Id            int             `json:"id"`
	TransactionId int             `json:"-"`
	ItemId        int             `json:"item_id"`
	Item          ItemTransaction `gorm:"ForeignKey:ItemId" json:"item"`
	Qty           int             `json:"qty"`
	Price         int             `json:"price"`
	Subtotal      int             `json:"subtotal"`
}

type AllTransactionItemResponse struct {
	Id            int `json:"id"`
	TransactionId int `json:"-"`
	Subtotal      int `json:"subtotal"`
}

func (TransactionItemResponse) TableName() string {
	return "transaction_items"
}

func (AllTransactionItemResponse) TableName() string {
	return "transaction_items"
}

// {
// "transaction_id": 1,
// "item_id": 4,
// "qty": 3,
// "price": 1000000,
// "subtotal": 1000000
// }
