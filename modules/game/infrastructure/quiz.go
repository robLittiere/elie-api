package infrastructure

import (
	"elie-api/modules/common/repository"
	"elie-api/modules/game/application/filters"
	"elie-api/modules/game/models"
	"gorm.io/gorm"
	"sort"
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

func (r *QuizRepo) FindQuizCompletedByUser(userId string) ([]int, error) {
	var quizIds []int
	_ = r.DB.Table("user_quizzes").Select("id").Where("user_id = ?", userId).Scan(&quizIds)
	return quizIds, nil
}

func (r *QuizRepo) FindNextQuiz(lastQuizID int) (*models.Quiz, error) {
	var allQuizzes []models.Quiz

	// Récupérer tous les quizzes
	r.DB = r.DB.Table("quiz_games").
		Select("quiz").
		Joins("JOIN jsonb_array_elements(data->'topic') as topic ON TRUE").
		Joins("JOIN jsonb_array_elements(topic->'quizzes') as quiz ON TRUE").
		Where("quiz->>'id' > ?", lastQuizID)

	// Exécuter la requête pour récupérer tous les quizzes
	result := r.DB.Find(&allQuizzes)
	if result.Error != nil {
		return nil, result.Error
	}

	// Trier les quizzes par ID
	sort.SliceStable(allQuizzes, func(i, j int) bool {
		return allQuizzes[i].Id < allQuizzes[j].Id
	})

	// Trouver le prochain quiz après lastQuizID
	var nextQuiz *models.Quiz
	for _, quiz := range allQuizzes {
		if quiz.Id > lastQuizID {
			nextQuiz = &quiz
			break
		}
	}

	// Si aucun quiz trouvé, retourner le premier quiz
	if nextQuiz == nil && len(allQuizzes) > 0 {
		nextQuiz = &allQuizzes[0]
	}

	return nextQuiz, nil
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
