package dto

type ModuleDTO struct {
	ID          string  `json:"id"`
	CourseID    string  `json:"course_id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	SortOrder   int     `json:"sort_order"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
	DeletedAt   *string `json:"deleted_at,omitempty"`
}
