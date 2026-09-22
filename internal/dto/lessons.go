package dto

type LessonDTO struct {
	ID              string   `json:"id"`
	ModuleID        string   `json:"module_id"`
	Title           string   `json:"title"`
	Slug            string   `json:"slug"`
	LessonType      string   `json:"lesson_type"`
	VideoSourceID   *string  `json:"video_source_id,omitempty"`
	DurationSeconds int      `json:"duration_seconds"`
	Content         *string  `json:"content,omitempty"`
	AttachmentURLs  []string `json:"attachment_urls,omitempty"`
	IsFreePreview   bool     `json:"is_free_preview"`
	SortOrder       int      `json:"sort_order"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	DeletedAt       *string  `json:"deleted_at,omitempty"`
}
