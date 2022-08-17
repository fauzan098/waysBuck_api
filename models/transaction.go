package models

type Transaction struct {
	ID     int                      `json:"id" gorm:"primary_key:auto_increment"`
	UserId int                      `json:"user_id"`
	User   UsersTransactionResponse `json:"user"`
	Cart   []Cart                   `json:"cart"`
	Amount int                      `json:"amount"`
}
