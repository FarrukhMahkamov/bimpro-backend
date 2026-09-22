package dto

type EnrollmentDTO struct {
	ID               string  `json:"id"`
	UserID           string  `json:"user_id"`
	CourseID         string  `json:"course_id"`
	IsActive         bool    `json:"is_active"`
	ExpiresAt        *string `json:"expires_at,omitempty"`
	GrantedBy        string  `json:"granted_by"`
	GrantedByAdminID *string `json:"granted_by_admin_id,omitempty"`
	CreatedAt        string  `json:"created_at"`
	DeletedAt        *string `json:"deleted_at,omitempty"`
}
