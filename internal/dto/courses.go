package dto

type CourseDTO struct {
	ID                   string   `json:"id"`
	AuthorID             string   `json:"author_id"`
	CategoryID           string   `json:"category_id"`
	Title                string   `json:"title"`
	Slug                 string   `json:"slug"`
	Subtitle             string   `json:"subtitle"`
	Description          string   `json:"description"`
	ThumbnailURL         string   `json:"thumbnail_url"`
	PromoVideoURL        string   `json:"promo_video_url"`
	Price                float64  `json:"price"`
	DiscountPrice        float64  `json:"discount_price"`
	Currency             string   `json:"currency"`
	TaxRate              float64  `json:"tax_rate"`
	Level                string   `json:"level"`
	Language             string   `json:"language"`
	Requirements         []string `json:"requirements"`
	WhatYouWillLearn     []string `json:"what_you_will_learn"`
	IsPublished          bool     `json:"is_published"`
	IsFeatured           bool     `json:"is_featured"`
	TotalDurationSeconds int      `json:"total_duration_seconds"`
	StudentsCount        int      `json:"students_count"`
	AverageRating        float64  `json:"average_rating"`
	CreatedAt            string   `json:"created_at"`
	UpdatedAt            string   `json:"updated_at"`
	DeletedAt            *string  `json:"deleted_at,omitempty"`
}
