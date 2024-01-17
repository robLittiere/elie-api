package models

import (
	userModel "elie-api/modules/user/models"
)

type UserQuiz struct {
	Id     int            `json:"id" gorm:"primary_key;auto_increment;not_null"`
	UserId int            `json:"user_id"`
	User   userModel.User `json:"user" gorm:"foreignKey:UserId"`
	QuizId int            `json:"quiz_id"`
}
