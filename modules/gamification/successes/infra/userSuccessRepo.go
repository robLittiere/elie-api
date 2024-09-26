package infra

import (
	"elie-api/modules/common/repository"
	models2 "elie-api/modules/gamification/successes/domain/models"
	"elie-api/modules/gamification/successes/infra/filters"
	"gorm.io/gorm"
)

type UserSuccessRepo struct {
	repository.BaseRepo
}

func NewUserSuccessRepo(db *gorm.DB) *UserSuccessRepo {
	return &UserSuccessRepo{BaseRepo: repository.BaseRepo{DB: db, FilterRegistry: filters.GetUserSuccessFilterRegitry()}}
}

func (repo *UserSuccessRepo) Find() ([]models2.UserSuccess, error) {
	userSuccess := make([]models2.UserSuccess, 0)
	result := repo.DB.Preload("Success.Tag").Find(&userSuccess)
	if result.Error != nil {
		return nil, result.Error
	}
	return userSuccess, nil
}

func (repo *UserSuccessRepo) BuildQueryAndFind(queryParams map[string][]string) ([]models2.UserSuccess, error) {
	var userSuccess []models2.UserSuccess
	err := repo.BuildQuery(queryParams)
	if err != nil {
		return nil, err
	}
	result := repo.DB.Preload("Success.Tag").Find(&userSuccess)
	if result.Error != nil {
		return nil, result.Error
	}

	return userSuccess, nil
}

func (r *UserSuccessRepo) FindByUuidAndUserSuccessId(uProgress *models2.UserSuccessProgressRequest, u *models2.UserSuccess) error {
	// Get user_id
	var uid int
	result := r.DB.Table("users").Select("id").Where("uuid = ?", uProgress.UserUuid).Scan(&uid)
	if result.Error != nil {
		return result.Error
	}

	result = r.DB.Preload("Success").Where("user_id = ? AND success_id = ?", uid, uProgress.SuccessId).First(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserSuccessRepo) FindUserSuccessByUserAndTag(tagId int, userId int, us *models2.UserSuccess) error {

	result := r.DB.Preload("Success").Joins("left join successes ON user_successes.success_id = successes.id").
		Where("user_successes.user_id = ? AND user_successes.is_completed = false AND successes.tag_id = ?", userId, tagId).
		First(&us)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserSuccessRepo) IncrementUserSuccessProgression(userSuccess *models2.UserSuccess) error {

	userSuccess.Progression++

	if userSuccess.Progression >= userSuccess.Success.DoneCondition {
		userSuccess.IsCompleted = true
	}

	result := r.DB.Save(&userSuccess)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserSuccessRepo) CreateUserSuccessFromUserAndSuccess(userId int, ns *models2.Success, us *models2.UserSuccess) error {

	newUserSuccess := &models2.UserSuccess{
		UserId:      userId,
		SuccessId:   ns.Id,
		Progression: us.Progression,
	}

	result := r.DB.Create(newUserSuccess)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
