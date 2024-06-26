package event

import (
	models2 "elie-api/modules/gamification/models"
	"elie-api/modules/user/models"
)

type QuestProgressListener interface {
	OnQuestProgress(user models.User, quest models2.Quest, userQuest models2.UserQuest) error
}
