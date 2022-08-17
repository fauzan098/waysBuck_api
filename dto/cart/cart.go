package cartdto

type CartRequest struct {
	ProductId     int `json:"product_id" gorm:"type: int"`
	ToppingID     int `json:"-" form:"topping_id" gorm:"-"`
	TransactionId int `json:"transaction_id" gorm:"type: int"`
	Qty           int `json:"qty" gorm:"type: int"`
	SubAmount     int `json:"sub_amount" gorm:"type: int"`
}
