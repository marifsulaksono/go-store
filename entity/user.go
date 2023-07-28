package entity

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Phonenumber int    `json:"phonenumber"`
	Role        string `json:"role"`
}

type UserResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"-"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Role        string `json:"role"`
}

type ShippingAddress struct {
	Id            int    `json:"id"`
	UserId        int    `json:"user_id"`
	RecipientName string `json:"recipient_name"`
	Address       string `json:"address"`
	Phonenumber   string `json:"phonenumber"`
}

func (UserResponse) TableName() string {
	return "users"
}
