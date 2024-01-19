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

func (r *QuizRepo) Find() ([]models.Quiz, error) {
	// var quizGame models.QuizGame
	var quizzes []models.Quiz

	r.DB = r.DB.Table("quiz_games").
		Select("quizzes").
		Joins("JOIN jsonb_array_elements(data->'topic') as topic ON TRUE").
		Joins("JOIN jsonb_array_elements(topic->'quizzes') as quizzes ON TRUE").
		Where("SELECT id FROM games WHERE name = ? AND game_version = ?", "Quiz", "1.0")

	// Add where clause to test
	r.DB = r.DB.Where("topic->>'name' = ?", "Energie Solaire")

	result := r.DB.Scan(&quizzes)

	if result.Error != nil {
		return quizzes, result.Error
	}
	return quizzes, nil
}

func (r *QuizRepo) BuildQueryAndFind(queryParams map[string][]string) ([]models.QuizGameJSONMap, error) {
	// var quizGame models.QuizGame
	quizzes := make([]models.QuizGameJSONMap, 0)

	r.DB = r.DB.Table("quiz_games").
		Select("quiz").
		Joins("JOIN jsonb_array_elements(data->'topic') as topic ON TRUE").
		Joins("JOIN jsonb_array_elements(topic->'quizzes') as quiz ON TRUE").
		Where("gid = ?", 4)

	err := r.BuildQuery(queryParams)
	if err != nil {
		return nil, err
	}

	// Execute query and get the results
	result := r.DB.Scan(&quizzes)
	if result.Error != nil {
		return nil, result.Error
	}

	return quizzes, nil
}
