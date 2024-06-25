package infrastructure

import (
	"elie-api/modules/common/repository"
	"elie-api/modules/game/application/filters"
	"elie-api/modules/game/models"
	"fmt"
	"gorm.io/gorm"
	"sort"
	"strconv"
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

func (r *QuizRepo) BuildQueryAndFindByData(id string, queryParams map[string][]string) ([]models.QuizGameJSONMap, error) {
	// var quizGame models.QuizGame
	quizzes := make([]models.QuizGameJSONMap, 0)

	r.DB = r.DB.Table("quiz_games").
		Select("quiz").
		Joins("JOIN jsonb_array_elements(data->'topic') as topic ON TRUE").
		Joins("JOIN jsonb_array_elements(topic->'quizzes') as quiz ON TRUE").
		Where("id = ?", id)

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

func (r *QuizRepo) CreateUserQuiz(quizGame *models.UserQuiz) error {
	result := r.DB.Create(quizGame)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *QuizRepo) FindQuizCompletedByUser(userId int) ([]string, error) {
	var quizIds []string
	_ = r.DB.Table("user_quizzes").Select("quiz_id").Where("user_id = ?", userId).Scan(&quizIds)
	return quizIds, nil
}

func (r *QuizRepo) FindNextQuiz(quizIds []string) *models.NextQuiz {
	var lastQuizID int

	if len(quizIds) > 0 {
		lastQuizID, _ = strconv.Atoi(quizIds[len(quizIds)-1])
	}

	var allQuizzes []models.NextQuiz
	fmt.Printf("Last Quiz ID: %d\n", lastQuizID)

	query := r.DB.Table("quiz_games").
		Select("quiz->>'id' as id, quiz->>'title' as title, quiz->>'topic' as topic, quiz->>'questions' as questions").
		Joins("JOIN jsonb_array_elements(data->'topic') as topic ON TRUE").
		Joins("JOIN jsonb_array_elements(topic->'quizzes') as quiz ON TRUE")

	if len(quizIds) > 0 {
		query = query.Where("quiz->>'id' NOT IN (?)", quizIds)
	}

	query.Scan(&allQuizzes)
	fmt.Printf("All Quizzes: %v\n", allQuizzes)

	sort.SliceStable(allQuizzes, func(i, j int) bool {
		return allQuizzes[i].Id < allQuizzes[j].Id
	})

	var nextQuiz *models.NextQuiz
	for _, quiz := range allQuizzes {
		if quiz.Id > lastQuizID {
			nextQuiz = &quiz
			break
		}
	}

	return nextQuiz
}

func (r *QuizRepo) GetRandomQuiz() (models.Quiz, error) {
	var quizData models.QuizGameJSONMap
	var quiz models.Quiz
	result := r.DB.Table("quiz_games").
		Select("quizData").
		Joins("JOIN jsonb_array_elements(data->'topic') as topic ON TRUE").
		Joins("JOIN jsonb_array_elements(topic->'quizzes') as quizData ON TRUE").
		Order("RANDOM()").
		Limit(1).
		Scan(&quizData)

	if result.Error != nil {
		return quiz, result.Error
	}

	err := quiz.LoadFromMap(quizData)
	if err != nil {
		return quiz, err
	}

	return quiz, nil
}
