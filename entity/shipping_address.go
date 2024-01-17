package entity

type ShippingAddress struct {
	Id            int    `gorm:"primaryKey,autoIncrement" json:"id"`
	UserId        int    `gorm:"not null" json:"user_id"`
	RecipientName string `gorm:"not null;size:255" json:"name"`
	Address       string `gorm:"not null" json:"address"`
	Phonenumber   string `gorm:"not null;size:16" json:"phonenumber"`
}

// {
// 	"recepient_name": "Arif",
// 	"Address": "Probolinggo",
// 	"phonenumber": "81234567890"
// }
