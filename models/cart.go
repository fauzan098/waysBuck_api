package models

type Cart struct {
	ID            int         `json:"id" gorm:"primary_key:auto_increment"`
	ProductId     int         `json:"product_id" gorm:"type: int"`
	Product       Product     `json:"product"`
	ToppingID     []int       `json:"-" form:"topping_id" gorm:"-"`
	Topping       []Topping   `json:"topping" gorm:"many2many:cart_toppings;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TransactionId int         `json:"transaction_id" gorm:"type: int"`
	Transaction   Transaction `json:"-"`
	Qty           int         `json:"qty" gorm:"type: int"`
	SubAmount     int         `json:"sub_amount" gorm:"type: int"`
}
