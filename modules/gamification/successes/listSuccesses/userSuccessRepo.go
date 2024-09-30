package listSuccesses

import (
	"elie-api/modules/common/repository"
	models2 "elie-api/modules/gamification/successes/domain/models"
	"elie-api/modules/gamification/successes/infra/filters"
	userModels "elie-api/modules/user/models"
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
	result := repo.DB.Preload("Success.Tag").Preload("Success.ParentSuccess").Find(&userSuccess)
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
	result := repo.DB.Preload("Success.Tag").Preload("Success.ParentSuccess").Find(&userSuccess)
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

func (r *UserSuccessRepo) AddCurrencyAmountSuccessToUser(user *userModels.User, userSuccess *models2.UserSuccess) error {
	user.CurrencyAmount += userSuccess.Success.CurrencyReward

	result := r.DB.Model(user).Updates(map[string]interface{}{
		"currency_amount": user.CurrencyAmount,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// TODO : Change this, this is really bad
func (r *UserSuccessRepo) FindTheNextProgressionRankSuccessIdByTag(us *models2.UserSuccess, ns *models2.Success) error {

	var count int64
	r.DB.Table("successes").
		Where("tag_id = ? AND progression_rank = ?", us.Success.TagId, us.Success.ProgressionRank+1).
		Count(&count)

	if count == 0 {
		// Success is already maxed
		return gorm.ErrRecordNotFound
	}

	result := r.DB.Table("successes").Preload("Tag").
		Where("tag_id = ? AND progression_rank = (SELECT progression_rank + 1 FROM successes WHERE id = ?)", us.Success.TagId, us.SuccessId).
		Pluck("id", &ns)
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
