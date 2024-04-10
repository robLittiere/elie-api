package infrastructure

import (
	"elie-api/modules/common/repository"
	"elie-api/modules/gamification/application/filters"
	"elie-api/modules/gamification/models"
	"gorm.io/gorm"
)

type SuccessRepo struct {
	repository.BaseRepo
}

func NewSuccessRepo(db *gorm.DB) *SuccessRepo {
	return &SuccessRepo{BaseRepo: repository.BaseRepo{DB: db, FilterRegistry: filters.GetSuccessFilterRegitry()}}
}

// Find Get all successes from db
func (r *SuccessRepo) Find() ([]models.Success, error) {
	success := make([]models.Success, 0)
	result := r.DB.Find(&success)
	if result.Error != nil {
		return nil, result.Error
	}
	return success, nil
}

func (r *SuccessRepo) GetByTag(tag string) ([]models.Success, error) {
	success := make([]models.Success, 0)
	result := r.DB.Where("tags = ?", tag).Find(&success)

	if result.Error != nil {
		return nil, result.Error
	}
	return success, nil
}
