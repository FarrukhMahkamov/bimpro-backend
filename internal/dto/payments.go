package dto

type PaymentDTO struct {
	ID                    string  `json:"id"`
	OrderID               string  `json:"order_id"`
	OrganizationID        *string `json:"organization_id,omitempty"`
	Provider              string  `json:"provider"`
	ProviderTransactionID *string `json:"provider_transaction_id,omitempty"`
	Amount                string  `json:"amount"`
	Currency              string  `json:"currency"`
	Status                string  `json:"status"`
	RawRequest            *string `json:"raw_request,omitempty"`
	RawResponse           *string `json:"raw_response,omitempty"`
	CreatedAt             string  `json:"created_at"`
}
