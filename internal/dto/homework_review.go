package dto

type HomeworkReviewDTO struct {
	ID           string  `json:"id"`
	SubmissionID string  `json:"submission_id"`
	InstructorID string  `json:"instructor_id"`
	ReviewText   string  `json:"review_text"`
	Rating       int     `json:"rating"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    *string `json:"deleted_at,omitempty"`
}
