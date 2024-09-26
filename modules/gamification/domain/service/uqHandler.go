package service

import (
	"elie-api/modules/gamification/domain/event"
	infrastructure2 "elie-api/modules/gamification/infrastructure"
	"elie-api/modules/gamification/models"
	"elie-api/modules/user/infrastructure"
	"gorm.io/gorm"
)

type UserQuestService struct {
	UserQuestRepo          *infrastructure2.UserQuestRepo
	UserRepo               *infrastructure.UserRepo
	QuestRepo              *infrastructure2.QuestRepo
	QuestProgressListeners []event.QuestProgressListener
}

func NewUserQuestService(conn *gorm.DB) *UserQuestService {
	return &UserQuestService{
		UserQuestRepo: infrastructure2.NewUserQuestRepo(conn),
		UserRepo:      infrastructure.NewUserRepo(conn),
		QuestRepo:     infrastructure2.NewQuestRepo(conn),
	}
}

func (s *UserQuestService) CreateUserQuest(userUuid string, qid int) error {
	user, err := s.UserRepo.FindByUuid(userUuid)
	if err != nil {
		return err
	}

	quest, err := s.QuestRepo.FindById(qid)
	if err != nil {
		return err
	}

	newUserQuest := &models.UserQuest{
		UserId:  user.Id,
		QuestId: quest.Id,
		Quest:   quest,
	}

	err = s.UserQuestRepo.Create(newUserQuest)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserQuestService) HandleUserQuestProgress(userUuid string, qid int) error {
	// TODO Return specifics errors for user or quest not found
	user, err := s.UserRepo.FindByUuid(userUuid)
	if err != nil {
		return err
	}
	quest, err := s.QuestRepo.FindById(qid)
	if err != nil {
		return err
	}
	userQuest, err := s.UserQuestRepo.FindByUserIdAndQid(user.Id, quest.Id)
	if err != nil {
		return err
	}

	// Increment progression and check if quest is completed
	userQuest.Progression += 1
	if userQuest.Progression >= userQuest.Quest.DoneCondition {
		userQuest.IsCompleted = true
	}
	err = s.UserQuestRepo.Update(&userQuest)

	// Update user xp
	if userQuest.IsCompleted {
		err = s.UserRepo.IncreaseUserXpQuest(&userQuest, &user)
		if err != nil {
			return err
		}
	}

	return nil
}
