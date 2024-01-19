package infrastructure

import (
	"elie-api/modules/common/repository"
	"elie-api/modules/gamification/application/filters"
	"elie-api/modules/gamification/models"
	"gorm.io/gorm"
)

type UserQuestRepo struct {
	repository.BaseRepo
}

func NewUserQuestRepo(db *gorm.DB) *UserQuestRepo {
	return &UserQuestRepo{BaseRepo: repository.BaseRepo{DB: db, FilterRegistry: filters.GetUserQuestFilterRegitry()}}
}

func (repo *UserQuestRepo) Find() ([]models.UserQuest, error) {
	userQuests := make([]models.UserQuest, 0)
	result := repo.DB.Preload("Quest").Find(&userQuests)
	if result.Error != nil {
		return nil, result.Error
	}
	return userQuests, nil
}

func (repo *UserQuestRepo) BuildQueryAndFind(queryParams map[string][]string) ([]models.UserQuest, error) {
	var userQuests []models.UserQuest
	err := repo.BuildQuery(queryParams)
	if err != nil {
		return nil, err
	}
	result := repo.DB.Preload("Quest").Find(&userQuests)
	if result.Error != nil {
		return nil, result.Error
	}

	return userQuests, nil
}
