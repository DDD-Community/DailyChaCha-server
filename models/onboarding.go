package models

type ExerciseGoal struct {
	UserID       int    `json:"user_id"`
	ExerciseGoal string `json:"exercise_goal"`
}

type ExerciseDate struct {
	UserID       int    `json:"user_id"`
	ExerciseDate string `json:"exercise_date"`
	ExerciseTime *int   `json:"exercise_time"`
}
