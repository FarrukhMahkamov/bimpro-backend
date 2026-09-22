package dto

type QuizAttemptDTO struct {
	ID         string  `json:"id"`
	QuizID     string  `json:"quiz_id"`
	UserID     string  `json:"user_id"`
	Score      int     `json:"score"`
	IsPassed   bool    `json:"is_passed"`
	AnswersLog *string `json:"answers_log,omitempty"`
	CreatedAt  string  `json:"created_at"`
}
