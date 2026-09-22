package dto

type UserDailyActivityDTO struct {
	UserID         string `json:"user_id"`
	ActivityDate   string `json:"activity_date"`
	LessonsWatched int    `json:"lessons_watched"`
	MinutesSpent   int    `json:"minutes_spent"`
}
