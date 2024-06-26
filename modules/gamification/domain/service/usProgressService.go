package service

import (
	error2 "elie-api/modules/gamification/domain/error"
	infrastructure2 "elie-api/modules/gamification/infrastructure"
	models2 "elie-api/modules/gamification/models"
	"elie-api/modules/user/infrastructure"
	"elie-api/modules/user/models"
	"errors"
	"gorm.io/gorm"
	"log"
)

type UserSuccessProgressService struct {
	UserRepo        *infrastructure.UserRepo
	UserSuccessRepo *infrastructure2.UserSuccessRepo
}

func NewUserSuccessProgressService(conn *gorm.DB) *UserSuccessProgressService {
	return &UserSuccessProgressService{
		UserRepo:        infrastructure.NewUserRepo(conn),
		UserSuccessRepo: infrastructure2.NewUserSuccessRepo(conn),
	}
}

func (u UserSuccessProgressService) OnQuestProgress(user models.User, quest models2.Quest, userQuest models2.UserQuest) error {
	if userQuest.IsCompleted {
		err := u.onQuestIsCompleted(user, quest, userQuest)
		if err != nil {
			if errors.As(err, &error2.SuccessNotFoundError{}) {
				log.Printf("UserSuccessProgressService.OnQuestProgress: %v", err.Error())
				return nil
			}
			if errors.As(err, &error2.NextLevelNotFoundError{}) {
				log.Printf("UserSuccessProgressService.OnQuestProgress: %v", err.Error())
				return nil
			}
			if errors.As(err, &error2.NextSuccessNotFound{}) {
				log.Printf("UserSuccessProgressService.OnQuestProgress: %v", err.Error())
				return nil
			}
			return err
		}
	}
	return nil
}

// TODO : This is really bad, we need to refactor this
// TODO : There is a lot of business logic in the repositories that are mixed with the queries
func (u UserSuccessProgressService) onQuestIsCompleted(user models.User, quest models2.Quest, userQuest models2.UserQuest) error {
	var userSuccess models2.UserSuccess
	if err := u.UserSuccessRepo.FindUserSuccessByUserAndTag(quest.TagId, user.Id, &userSuccess); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return error2.SuccessNotFoundError{
				Args: map[string]interface{}{
					"tagId":  quest.TagId,
					"userId": user.Id,
				},
			}
		}
		return err
	}

	if err := u.UserSuccessRepo.IncrementUserSuccessProgression(&userSuccess); err != nil {
		return err
	}

	if userSuccess.IsCompleted {
		if err := u.UserRepo.IncreaseUserXpSuccess(&userSuccess, &user); err != nil {
			return err
		}

		if err := u.UserSuccessRepo.AddCurrencyAmountSuccessToUser(&user, &userSuccess); err != nil {
			return err
		}

		var nextSuccess models2.Success
		if err := u.UserSuccessRepo.FindTheNextProgressionRankSuccessIdByTag(&userSuccess, &nextSuccess); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return error2.NextSuccessNotFound{
					Args: map[string]interface{}{
						"userSuccess": userSuccess,
					},
				}
			}
			return err
		}

		if err := u.UserSuccessRepo.CreateUserSuccessFromUserAndSuccess(user.Id, &nextSuccess, &userSuccess); err != nil {
			return err
		}
	}
	return nil
}
