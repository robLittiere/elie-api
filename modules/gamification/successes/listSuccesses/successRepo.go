package listSuccesses

import (
	"elie-api/modules/common/repository"
	"elie-api/modules/gamification/successes/domain/models"
	"elie-api/modules/gamification/successes/infra/filters"
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
	result := r.DB.Preload("Tag").Preload("ParentSuccess").Find(&success)
	if result.Error != nil {
		return nil, result.Error
	}
	return success, nil
}

func (r *SuccessRepo) BuildQueryAndFind(queryParams map[string][]string) ([]models.Success, error) {
	var successes []models.Success
	err := r.BuildQuery(queryParams)
	if err != nil {
		return nil, err
	}
	successes, err = r.Find()
	if err != nil {
		return nil, err
	}

	return successes, nil
}
