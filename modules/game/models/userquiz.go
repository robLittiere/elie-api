package models

type UserQuiz struct {
	Id     int    `json:"id" gorm:"primary_key;auto_increment;not_null"`
	UserID int    `json:"user_id"`
	QuizId string `json:"quiz_id"`
}

type UserQuizRequest struct {
	UserUuid string `json:"user_uuid"`
	QuizId   string `json:"quiz_id"`
}
