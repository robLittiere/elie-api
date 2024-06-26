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

func (repo *UserQuestRepo) FindOne() (models.UserQuest, error) {
	var userQuest models.UserQuest
	result := repo.DB.Preload("Quest").First(&userQuest)
	if result.Error != nil {
		return userQuest, result.Error
	}
	return userQuest, nil
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

func (r *UserQuestRepo) Create(u *models.UserQuest) error {
	result := r.DB.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserQuestRepo) Update(u *models.UserQuest) error {
	result := r.DB.Save(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserQuestRepo) IncrementUserQuestProgression(userQuest *models.UserQuest) error {

	userQuest.Progression += 1

	if userQuest.Progression >= userQuest.Quest.DoneCondition {
		userQuest.IsCompleted = true
	}

	result := r.DB.Model(userQuest).Updates(map[string]interface{}{
		"progression":  userQuest.Progression,
		"is_completed": userQuest.IsCompleted,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
