package dto

type QuizDTO struct {
	ID           string `json:"id"`
	LessonID     string `json:"lesson_id"`
	Title        string `json:"title"`
	PassingScore int    `json:"passing_score"`
	CreatedAt    string `json:"created_at"`
}
