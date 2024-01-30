package entity

type PaymentRequestPayload struct {
	TransactionDetails TransactionDetail `json:"transaction_details"`
	ItemDetail         []ItemDetails     `json:"item_details"`
	Expiry             *ExpiryDetails    `json:"expiry,omitempty"`
	CustomerDetail     *CostumerDetails  `json:"customer_details,omitempty"`
}

type TransactionDetail struct {
	OrderId  string `json:"order_id"`
	GrossAmt int    `json:"gross_amount"`
}

type ItemDetails struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Qty          int    `json:"quantity"`
	MerchantName string `json:"merchant_name,omitempty"`
}

type CostumerDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type ExpiryDetails struct {
	Start    string `json:"start_time"`
	Duration int    `json:"duration"`
	Unit     string `json:"unit"`
}
