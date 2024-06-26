package infrastructure

import (
	"elie-api/modules/common/repository"
	error2 "elie-api/modules/gamification/domain/error"
	models2 "elie-api/modules/gamification/models"
	"elie-api/modules/user/application/filters"
	"elie-api/modules/user/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	repository.BaseRepo
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	// Get the available filters for this repo
	return &UserRepo{BaseRepo: repository.BaseRepo{DB: db, FilterRegistry: filters.GetUserFilters()}}
}

func (r *UserRepo) Find() ([]models.User, error) {
	var users []models.User
	result := r.DB.Preload("Level").Preload("UserQuests.Quest").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepo) BuildQueryAndFind(queryParams map[string][]string) ([]models.User, error) {
	var users []models.User
	err := r.BuildQuery(queryParams)
	if err != nil {
		return nil, err
	}
	users, err = r.Find()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) CreateUser(user *models.User) error {
	if err := r.IsUserInDb(*user); err != nil {
		return err
	}

	user.Uuid = uuid.New().String()

	// Get basic level and add it to the user
	// TODO improve this by adding a user builder or something as we will need to add more default stuff to the user
	var level models2.Level
	result := r.DB.Table("levels").Select("id").Where("name = ?", "beginner").Scan(&level)
	if result.Error != nil {
		return result.Error
	}
	user.LevelId = level.Id

	if err := user.HashPassword(user.Password, 2); err != nil {
		return err
	}

	result = r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	// Assign all level 1 successes to the user
	var successes []models2.Success
	if err := r.DB.Table("successes").Select("id").Where("progression_rank = ?", 1).Find(&successes).Error; err != nil {
		return err
	}

	for _, success := range successes {
		userSuccess := models2.UserSuccess{
			UserId:    user.Id,
			SuccessId: success.Id,
		}
		if err := r.DB.Create(&userSuccess).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *UserRepo) IsUserInDb(user models.User) error {
	result := r.DB.Select("id").Where("email = ?", user.Email).Limit(1).Find(&models.User{})
	if result.RowsAffected != 0 {
		err := fmt.Errorf("user already registered with this email")
		return err
	}
	return nil
}

func (r *UserRepo) FindByEmail(email string) (models.User, error) {
	var user models.User
	result := r.DB.Preload("Level").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil

}

func (r *UserRepo) FindByUuid(uid string) (models.User, error) {
	var user models.User
	result := r.DB.Preload("Level").Where("uuid = ?", uid).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (r *UserRepo) UpdateUser(user *models.User) error {
	var existingUser models.User
	result := r.DB.Select("id").Where("uuid = ?", user.Uuid).First(&existingUser)
	if result.Error != nil {
		return result.Error
	}

	result = r.DB.Model(&existingUser).Updates(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// TODO : WARNIIING !!!! Refactor this is really bad
// TODO : There is business logic in the infrastructure layer
func (r *UserRepo) IncreaseUserXpQuest(userQuest *models2.UserQuest, user *models.User) error {

	user.Xp += userQuest.Quest.Xp
	// Check if the user has enough xp to level up
	if user.Xp >= user.Level.NextLevelXpRequirement {
		var nextLevel models2.Level
		result := r.DB.Where("level_number", user.Level.LevelNumber+1).First(&nextLevel)

		// We need to get the next level in order to know if the user is already max level
		if result.Error != nil {
			// User is already max level as there is no next level
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}

		// Increase user currency
		err := r.IncreaseUserCurrency(user, user.Level.CurrencyWon)
		if err != nil {
			return err
		}

		user.LevelId = nextLevel.Id
		// Update user xp from previous level
		// This way leftover xp is not lost
		// For example, if the user has 20 xp but level required 10 xp to be passed
		// He would have 'overlevelup' of 10 xp. This way we carry those 10 xp on the next level
		user.Xp -= user.Level.NextLevelXpRequirement
		user.Level = nextLevel

	}

	result := r.DB.Model(&user).Updates(map[string]interface{}{
		"xp":       user.Xp,
		"level_id": user.LevelId,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// TODO : WARNIIING !!!! Refactor this is really bad
// TODO : There is business logic in the infrastructure layer
func (r *UserRepo) IncreaseUserXpSuccess(userSuccess *models2.UserSuccess, user *models.User) error {

	user.Xp += userSuccess.Success.Xp
	// Check if the user has enough xp to level up
	if user.Xp >= user.Level.NextLevelXpRequirement {

		var nextLevel models2.Level
		result := r.DB.Where("level_number", user.Level.LevelNumber+1).First(&nextLevel)

		// We need to get the next level in order to know if the user is already max level
		if result.Error != nil {
			// User is already max level as there is no next level
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return error2.NextLevelNotFoundError{
					Args: map[string]interface{}{
						"MaxLevelId": user.LevelId,
					},
				}
			}
			return result.Error
		}

		// Increase user currency
		err := r.IncreaseUserCurrency(user, user.Level.CurrencyWon)
		if err != nil {
			return err
		}

		user.LevelId = nextLevel.Id
		// Update user xp from previous level
		// This way leftover xp is not lost
		// For example, if the user has 20 xp but level required 10 xp to be passed
		// He would have 'overlevelup' of 10 xp. This way we carry those 10 xp on the next level
		user.Xp -= user.Level.NextLevelXpRequirement
		user.Level = nextLevel
	}

	result := r.DB.Model(&user).Updates(map[string]interface{}{
		"xp":       user.Xp,
		"level_id": user.LevelId,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) IncreaseUserCurrency(user *models.User, amount int) error {
	user.CurrencyAmount += amount

	result := r.DB.Model(user).Updates(map[string]interface{}{
		"currency_amount": user.CurrencyAmount,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
