package infrastructure

import (
	"elie-api/modules/common/repository"
	"elie-api/modules/game/application/filters"
	"elie-api/modules/game/models"
	"gorm.io/gorm"
)

type QuizRepo struct {
	repository.BaseRepo
}

func NewQuizRepo(db *gorm.DB) *QuizRepo {
	return &QuizRepo{BaseRepo: repository.BaseRepo{DB: db, FilterRegistry: filters.GetQuizFilterRegistry()}}
}

func (repo *QuizRepo) Find() (models.QuizGame, error) {
	var quizGame models.QuizGame

	result := repo.DB.Where("gid = ?", 2).First(&quizGame)
	if result.Error != nil {
		return quizGame, result.Error
	}
	return quizGame, nil
}
