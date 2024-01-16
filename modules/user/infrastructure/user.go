package infrastructure

import (
	"elie-api/modules/common"
	"elie-api/modules/user/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) Find() ([]models.User, error) {
	var users []models.User
	result := r.DB.Preload("Level").Preload("Quests").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepo) ApplyCriteria(c common.Criteria, queryValue string) {
	r.DB = c.ApplyQuery(r.DB, queryValue)
}
