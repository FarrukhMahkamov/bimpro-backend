package dto

// CREATE TABLE users (
//     id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

//     email              VARCHAR(255) UNIQUE,
//     phone              VARCHAR(20) UNIQUE NOT NULL,
//     is_phone_verified  BOOLEAN DEFAULT FALSE,
//     is_email_verified  BOOLEAN DEFAULT FALSE,

//     password_hash        VARCHAR(255) NOT NULL,
//     password_updated_at  TIMESTAMPTZ DEFAULT NOW(),
//     must_change_password BOOLEAN DEFAULT FALSE,
//     token_version        INT DEFAULT 1,

//     first_name  VARCHAR(100) NOT NULL,
//     last_name   VARCHAR(100) NOT NULL,
//     middle_name VARCHAR(100),
//     avatar_url  TEXT,

//     is_active  BOOLEAN DEFAULT TRUE,
//     is_banned  BOOLEAN DEFAULT FALSE,
//     ban_reason TEXT,

//     specialization   VARCHAR(100),
//     company_name     VARCHAR(150),
//     experience_years INT,

//     signup_ip    VARCHAR(45),
//     utm_source   VARCHAR(100),
//     utm_medium   VARCHAR(100),
//     utm_campaign VARCHAR(100),
//     referred_by  UUID REFERENCES users(id) ON DELETE SET NULL,

//     created_at    TIMESTAMPTZ DEFAULT NOW(),
//     updated_at    TIMESTAMPTZ DEFAULT NOW(),
//     last_login_at TIMESTAMPTZ
// );

//UserDTO is a Data Transfer Object for the User model.
type UserDTO struct {
	ID                uint64 `json:"id"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	IsPhoneVerified   bool   `json:"is_phone_verified"`
	IsEmailVerified   bool   `json:"is_email_verified"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	MiddleName        string `json:"middle_name,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	PasswordHash      string `json:"password_hash,omitempty"`
	PasswordUpdatedAt string `json:"password_updated_at,omitempty"`
	TokenVersion      int    `json:"token_version,omitempty"`
	SignupIP          string `json:"signup_ip,omitempty"`

	// Account status
	IsActive  bool   `json:"is_active"`
	IsBanned  bool   `json:"is_banned"`
	BanReason string `json:"ban_reason,omitempty"`

	// Professional details
	Specialization  string `json:"specialization,omitempty"`
	CompanyName     string `json:"company_name,omitempty"`
	ExperienceYears int    `json:"experience_years,omitempty"`

	// Timestamps
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	LastLoginAt string `json:"last_login_at,omitempty"`
}
