package dto

// UserDTO представляет собой структуру данных для передачи информации о пользователе между слоями приложения.
type UserDTO struct {
	ID              string  `json:"id"`
	Email           *string `json:"email,omitempty"` // Указатель, так как в базе email может быть NULL (UNIQUE, но nullable)
	Phone           string  `json:"phone"`
	IsPhoneVerified bool    `json:"is_phone_verified"`
	IsEmailVerified bool    `json:"is_email_verified"`

	// Безопасность и сессии (обычно не отдаются наружу все вместе, но если это внутренний DTO — ок)
	PasswordHash       string `json:"password_hash,omitempty"`
	PasswordUpdatedAt  string `json:"password_updated_at,omitempty"`
	MustChangePassword bool   `json:"must_change_password"`
	TokenVersion       int    `json:"token_version,omitempty"`

	// Профиль
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	MiddleName *string `json:"middle_name,omitempty"` // Использование *string удобно для корректной работы с NULL из БД
	AvatarURL  *string `json:"avatar_url,omitempty"`

	// Статус
	IsActive  bool    `json:"is_active"`
	IsBanned  bool    `json:"is_banned"`
	BanReason *string `json:"ban_reason,omitempty"`

	// Профессиональные данные
	Specialization  *string `json:"specialization,omitempty"`
	CompanyName     *string `json:"company_name,omitempty"`
	ExperienceYears *int    `json:"experience_years,omitempty"`

	// Маркетинг и аналитика (из вашего предыдущего вопроса)
	SignupIP    *string `json:"signup_ip,omitempty"`
	UtmSource   *string `json:"utm_source,omitempty"`
	UtmMedium   *string `json:"utm_medium,omitempty"`
	UtmCampaign *string `json:"utm_campaign,omitempty"`
	ReferredBy  *string `json:"referred_by,omitempty"` // UUID пригласившего

	// Таймстампы
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	LastLoginAt *string `json:"last_login_at,omitempty"`
}
