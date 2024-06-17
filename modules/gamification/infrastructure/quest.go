package infrastructure

import (
	"elie-api/modules/common/repository"
	"elie-api/modules/gamification/application/filters"
	"elie-api/modules/gamification/models"
	"gorm.io/gorm"
)

type QuestRepo struct {
	repository.BaseRepo
}

func NewQuestRepo(db *gorm.DB) *QuestRepo {
	return &QuestRepo{BaseRepo: repository.BaseRepo{DB: db, FilterRegistry: filters.GetQuestFilterRegistry()}}
}

func (repo *QuestRepo) Find() ([]models.Quest, error) {
	quests := make([]models.Quest, 0)
	result := repo.DB.Find(&quests)
	if result.Error != nil {
		return nil, result.Error
	}
	return quests, nil
}

func (repo *QuestRepo) BuildQueryAndFind(queryParams map[string][]string) ([]models.Quest, error) {
	var quests []models.Quest
	err := repo.BuildQuery(queryParams)
	if err != nil {
		return nil, err
	}
	quests, err = repo.Find()
	if err != nil {
		return nil, err
	}

	return quests, nil
}
