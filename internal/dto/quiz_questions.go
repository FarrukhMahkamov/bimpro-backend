package dto

type QuizQuestionDTO struct {
	ID                   string   `json:"id"`
	QuizID               string   `json:"quiz_id"`
	QuestionText         string   `json:"question_text"`
	Options              []string `json:"options"`
	CorrectAnswerIndices []int    `json:"correct_answer_indices"`
	SortOrder            int      `json:"sort_order"`
	CreatedAt            string   `json:"created_at"`
	UpdatedAt            string   `json:"updated_at"`
	DeletedAt            *string  `json:"deleted_at,omitempty"`
}
