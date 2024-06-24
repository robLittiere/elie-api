package tests

import (
	"elie-api/modules/fixtures"
	gameFixtures "elie-api/modules/fixtures/gamification"
	userFixtures "elie-api/modules/fixtures/user"
	gamificationController "elie-api/modules/gamification/application/controllers"
	gamificationModels "elie-api/modules/gamification/models"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"testing"
)

func TestIShouldProgressQuestWithProgressHandler(t *testing.T) {
	c, w = CreateGinTestContext()
	data := map[string]interface{}{
		"DoneCondition": 3,
	}
	quest := gameFixtures.CreateQuest(data)

	user := userFixtures.CreateUser(map[string]interface{}{})
	gameFixtures.CreateUserQuest(user, quest)

	uid, _ := uuid.Parse(user.Uuid)
	uq := gamificationModels.UserQuestProgressRequest{
		UserUuid:    uid,
		QuestId: quest.Id,
	}

	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUserQuest := gameFixtures.GetUserQuest(quest.Id, user.Id)

	assert.Equal(t, newUserQuest.Progression, 1)

	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUserQuest = gameFixtures.GetUserQuest(quest.Id, user.Id)

	assert.Equal(t, newUserQuest.Progression, 2)

	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUserQuest = gameFixtures.GetUserQuest(quest.Id, user.Id)

	assert.Equal(t, newUserQuest.Progression, 3)
	assert.Equal(t, newUserQuest.IsCompleted, true)
}

func TestQuestShouldNotBeCompletedIfProgressionIsNotMet(t *testing.T) {
	c, w = CreateGinTestContext()
	data := map[string]interface{}{
		"DoneCondition": 3,
	}
	quest := gameFixtures.CreateQuest(data)

	user := userFixtures.CreateUser(map[string]interface{}{})
	gameFixtures.CreateUserQuest(user, quest)

	uid, _ := uuid.Parse(user.Uuid)
	uq := gamificationModels.UserQuestProgressRequest{
		UserUuid:    uid,
		QuestId: quest.Id,
	}

	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUserQuest := gameFixtures.GetUserQuest(quest.Id, user.Id)

	assert.Equal(t, newUserQuest.Progression, 1)

	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUserQuest = gameFixtures.GetUserQuest(quest.Id, user.Id)

	assert.Equal(t, newUserQuest.Progression, 2)
	if newUserQuest.IsCompleted {
		t.Errorf("Quest should not be completed")
	}
	assert.Equal(t, newUserQuest.IsCompleted, false)
}

func TestUserShouldGetExperienceFromCompletedQuest(t *testing.T) {
	c, w = CreateGinTestContext()
	// Introduce second level so user is not max level
	gameFixtures.CreateLevel(map[string]interface{}{
		"XpRequirement": 100,
		"LevelNumber":   2,
	})

	quest := gameFixtures.CreateQuest(map[string]interface{}{
		"DoneCondition": 2,
		"Xp":            5,
	})

	user := userFixtures.CreateUser(map[string]interface{}{
		"Xp": 0,
	})

	gameFixtures.CreateUserQuest(user, quest)

	uid, _ := uuid.Parse(user.Uuid)
	uq := gamificationModels.UserQuestProgressRequest{
		UserUuid:    uid,
		QuestId: quest.Id,
	}

	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUser := userFixtures.GetUserByID(user.Id)

	if newUser.Xp != 0 {
		t.Errorf("Quest is not completed and should not give xp")
	}
	assert.Equal(t, newUser.Xp, 0)

	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUser = userFixtures.GetUserByID(user.Id)

	if newUser.Xp != 5 {
		t.Errorf("Quest is completed and should give xp")
	}
	assert.Equal(t, newUser.Xp, 5)
}
