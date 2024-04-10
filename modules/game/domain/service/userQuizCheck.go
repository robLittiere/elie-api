package service

import (
	"elie-api/modules/game/infrastructure"
	"elie-api/modules/game/models"
	"fmt"
	"gorm.io/gorm"
)

type UserQuizService struct {
	QuizRepo *infrastructure.QuizRepo
}

func (s *UserQuizService) CheckUserQuizExist(userQuiz models.UserQuiz) (bool, error) {
	var existQuiz int64
	var quizGameId int
	result := s.QuizRepo.DB.Table("user_quizzes").Where("user_id = ? AND quiz_id = ?", userQuiz.UserID, userQuiz.QuizId).Count(&existQuiz)
	fmt.Printf("quizGameId: %v\n", quizGameId)

	if result.Error != nil {
		return true, result.Error
	}
	return false, nil
}

func NewUserQuizService(db *gorm.DB) *UserQuizService {
	return &UserQuizService{
		QuizRepo: infrastructure.NewQuizRepo(db),
	}
}
