package dto

type InstructorProfileDTO struct {
	UserID             string   `json:"user_id"`
	Title              string   `json:"title"`
	Bio                *string  `json:"bio,omitempty"`
	Rating             float64  `json:"rating"`
	TotalStudentsCount int      `json:"total_students_count"`
	SocialMediaLinks   []string `json:"social_media_links"`
	CreatedAt          string   `json:"created_at"`
	UpdatedAt          string   `json:"updated_at"`
	DeletedAt          *string  `json:"deleted_at,omitempty"`
}
