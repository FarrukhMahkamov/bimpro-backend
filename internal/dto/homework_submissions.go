package dto

type HomeworkSubmissionDTO struct {
	ID          string   `json:"id"`
	LessonID    string   `json:"lesson_id"`
	UserID      string   `json:"user_id"`
	FileURLs    []string `json:"file_urls"`
	CommentText *string  `json:"comment_text,omitempty"`
	Status      string   `json:"status"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}
