package tests

import (
	"elie-api/modules/gamification/domain/quest"
	models2 "elie-api/modules/gamification/models"
	"elie-api/modules/user/models"
	"testing"
)

func TestUserQuestShouldProgress(t *testing.T) {
	// Setup
	var user models.User
	userQuest := models2.UserQuest{
		Id:          1,
		UserId:      1,
		QuestId:     1,
		Quest:       models2.Quest{Id: 1, QuestTypeId: 1, Name: "test", Xp: 1, Difficulty: "test", DoneCondition: 3},
		Progression: 0,
		IsCompleted: false,
	}
	user.UserQuests = append(user.UserQuests, userQuest)

	// Domain logic
	quest.ProgressQuest(&user, &userQuest)

	// Assert
	if user.UserQuests[0].Progression != 1 {
		t.Errorf("Expected progression to be 1 but got %v", userQuest.Progression)
	}

	// Progress two times
	quest.ProgressQuest(&user, &userQuest)
	quest.ProgressQuest(&user, &userQuest)

	if user.UserQuests[0].Progression != 3 {
		t.Errorf("Expected progression to be 2 but got %v", userQuest.Progression)
	}
	if user.UserQuests[0].IsCompleted != true {
		t.Errorf("Expected quest to be completed but got %v", userQuest.IsCompleted)
	}
}
