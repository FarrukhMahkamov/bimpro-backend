package dto

// UserSessionDTO представляет собой структуру данных для передачи информации о сессии пользователя между слоями приложения.
type UserSessionDTO struct {
	ID               string  `json:"id"`
	UserID           string  `json:"user_id"`
	RefreshTokenHash string  `json:"refresh_token_hash"`
	RotatedAt        *string `json:"rotated_at,omitempty"`
	UserAgent        *string `json:"user_agent,omitempty"`
	IPAddress        *string `json:"ip_address,omitempty"`
	ExpiresAt        string  `json:"expires_at"`
	CreatedAt        string  `json:"created_at"`
}
