package dto

type LessonProgressDTO struct {
	ID                  string  `json:"id"`
	UserID              string  `json:"user_id"`
	LessonID            string  `json:"lesson_id"`
	IsCompleted         bool    `json:"is_completed"`
	WatchedSeconds      int     `json:"watched_seconds"`
	LastWatchedPosition int     `json:"last_watched_position"`
	CompletionRate      float64 `json:"completion_rate"`
	CompletedAt         *string `json:"completed_at,omitempty"`
	UpdatedAt           string  `json:"updated_at"`
}
