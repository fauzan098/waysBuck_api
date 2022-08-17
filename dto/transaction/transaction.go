package transactiondto

type TransactionRequest struct {
	UserID int `json:"user_id" gorm:"type: int"`
	Amount int `json:"amount" gorm:"type: int"`
}
