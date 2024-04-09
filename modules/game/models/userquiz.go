package models

import (
	userModel "elie-api/modules/user/models"
)

type UserQuiz struct {
	Id     int            `json:"id" gorm:"primary_key;auto_increment;not_null"`
	UserID int            `json:"user_id"`
	User   userModel.User `json:"user" gorm:"foreignKey:user_id"`
	QuizId string         `json:"quiz_id"`
}

type UserQuizRequest struct {
	UserUuid string `json:"user_uuid"`
	QuizId   string `json:"quiz_id"`
}
