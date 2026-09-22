package dto

type OrderDTO struct {
	ID             string  `json:"id"`
	UserID         string  `json:"user_id"`
	CourseID       *string `json:"course_id,omitempty"`
	BundleID       *string `json:"bundle_id,omitempty"`
	PromoCodeID    *string `json:"promo_code_id,omitempty"`
	Amount         string  `json:"amount"`
	DiscountAmount string  `json:"discount_amount"`
	FinalAmount    string  `json:"final_amount"`
	Status         string  `json:"status"`
	Currency       string  `json:"currency"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}
