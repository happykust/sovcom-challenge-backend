package payments

type PaymentsTransactionsGet struct {
	TransactionUUID string `json:"transaction_uuid"`
}

type PaymentsTransactionsGetResponse struct {
	TransactionUUID   string  `json:"TransactionUUID"`
	UserId            uint    `json:"UserId"`
	Amount            float64 `json:"Amount"`
	PromoCode         string  `json:"PromoCode" gorm:"default:'NO_PROMO'"`
	TransactionStatus string  `json:"TransactionStatus" gorm:"default:PROCESSING"`
}
