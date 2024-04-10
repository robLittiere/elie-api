package service

import (
	"elie-api/config"
	modelUserQuizz "elie-api/modules/game/models"
	infraUser "elie-api/modules/user/infrastructure"
	"fmt"
	"gorm.io/gorm"
)

func GetUserQuizFromRequest(db *gorm.DB, userquizRequest modelUserQuizz.UserQuizRequest) modelUserQuizz.UserQuiz {
	var userquiz modelUserQuizz.UserQuiz

	userRepo := infraUser.NewUserRepo(config.DB)
	user, err := userRepo.FindByUuid(userquizRequest.UserUuid)
	fmt.Printf("user: %v\n", user)
	if err != nil {
		return userquiz
	}

	userquiz.UserID = user.Id
	userquiz.QuizId = userquizRequest.QuizId

	return userquiz
}
