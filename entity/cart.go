package entity

type Cart struct {
	Id        int                        `json:"id"`
	UserId    int                        `json:"-"`
	ProductId *int                       `json:"product_id"`
	Product   ProductTransactionResponse `json:"product"`
	Qty       *int                       `json:"qty"`
}

/*

Body Request:
"POST":
{
	"product_id": 1,
	"qty": 10
}

"PUT":
{
	"qty": 20
}

*/
