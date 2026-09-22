package dto

type PromoCodeDTO struct {
	ID              string  `json:"id"`
	Code            string  `json:"code"`
	DiscountPercent *int    `json:"discount_percent,omitempty"`
	DiscountAmount  *string `json:"discount_amount,omitempty"`
	MaxUses         *int    `json:"max_uses,omitempty"`
	UsedCount       int     `json:"used_count"`
	ExpiresAt       *string `json:"expires_at,omitempty"`
	IsActive        bool    `json:"is_active"`
	CreatedAt       string  `json:"created_at"`
}
