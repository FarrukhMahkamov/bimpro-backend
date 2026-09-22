package dto

type RefundDTO struct {
	id        string  `json:"id"`
	UserID    string  `json:"user_id"`
	PaymentID string  `json:"payment_id"`
	Amount    string  `json:"amount"`
	Reason    *string `json:"reason,omitempty"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
}
