package entity

type Cart struct {
	Id        int                        `json:"id"`
	UserId    int                        `json:"-"`
	ProductId int                        `json:"product_id"`
	Product   ProductTransactionResponse `json:"product"`
	Qty       int                        `json:"qty"`
}
