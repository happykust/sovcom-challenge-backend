package payments

type PaymentsTransactionsGet struct {
	TransactionUUID string `json:"transaction_uuid"`
}

type PaymentsTransactionsGetResponse struct {
	TransactionUUID   string `json:"transaction_uuid"`
	UserId            uint   `json:"user_id"`
	Amount            string `json:"amount"`
	PromoCode         string `json:"promocode" gorm:"default:'NO_PROMO'"`
	TransactionStatus string `json:"transaction_status" gorm:"default:PROCESSING"`
}
