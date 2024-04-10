package infrastructure

import (
	"elie-api/modules/common/repository"
	"elie-api/modules/gamification/application/filters"
	"elie-api/modules/gamification/models"
	modelsUser "elie-api/modules/user/models"
	"gorm.io/gorm"
)

type UserSuccessRepo struct {
	repository.BaseRepo
}

func NewUserSuccessRepo(db *gorm.DB) *UserSuccessRepo {
	return &UserSuccessRepo{BaseRepo: repository.BaseRepo{DB: db, FilterRegistry: filters.GetUserSuccessFilterRegitry()}}
}

func (repo *UserSuccessRepo) Find() ([]models.UserSuccess, error) {
	userSuccess := make([]models.UserSuccess, 0)
	result := repo.DB.Preload("Success").Find(&userSuccess)
	if result.Error != nil {
		return nil, result.Error
	}
	return userSuccess, nil
}

func (repo *UserSuccessRepo) BuildQueryAndFind(queryParams map[string][]string) ([]models.UserSuccess, error) {
	var userSuccess []models.UserSuccess
	err := repo.BuildQuery(queryParams)
	if err != nil {
		return nil, err
	}
	result := repo.DB.Preload("Success").Find(&userSuccess)
	if result.Error != nil {
		return nil, result.Error
	}

	return userSuccess, nil
}

func (r *UserSuccessRepo) FindByUuidAndUserSuccessId(uProgress *models.UserSuccessProgressRequest, u *models.UserSuccess) error {
	// Get user_id
	var uid int
	result := r.DB.Table("users").Select("id").Where("uuid = ?", uProgress.UserUuid).Scan(&uid)
	if result.Error != nil {
		return result.Error
	}

	result = r.DB.Preload("Success").Where("user_id = ? AND success_id = ?", uid, uProgress.UserSuccessId).First(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserSuccessRepo) AddCurrencyAmountSuccessToUser(user *modelsUser.User, userSuccess *models.UserSuccess) error {
	user.CurrencyAmount += userSuccess.Success.CurrencyReward

	result := r.DB.Model(user).Updates(map[string]interface{}{
		"currency_amount": user.CurrencyAmount,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
