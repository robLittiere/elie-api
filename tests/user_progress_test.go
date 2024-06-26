package tests

import (
	"elie-api/modules/fixtures"
	"elie-api/modules/fixtures/common"
	gameFixtures "elie-api/modules/fixtures/gamification"
	userFixtures "elie-api/modules/fixtures/user"
	gamificationController "elie-api/modules/gamification/application/controllers"
	gamificationModels "elie-api/modules/gamification/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserProgressTestSuite struct {
	suite.Suite
}

func (t *UserProgressTestSuite) SetupTest() {
	Init()
	common.CreateBatchLevels()
}

func TestUserProgressTestSuite(t *testing.T) {
	suite.Run(t, new(UserProgressTestSuite))
}

func (t *UserProgressTestSuite) TestUserShouldGetCurrencyWhenLevelUp() {
	c, w = CreateGinTestContext()

	tag := gameFixtures.CreateTag(map[string]interface{}{
		"Name": gamificationModels.PlayQuizTag,
	})
	quest := gameFixtures.CreateQuest(map[string]interface{}{
		"TagId":         tag.Id,
		"Xp":            10,
		"DoneCondition": 1,
	})
	userLevel := gameFixtures.CreateLevel(map[string]interface{}{
		"NextLevelXpRequirement": 10,
		"LevelNumber":            10000,
		"CurrencyWon":            200,
	})
	nextLevel := gameFixtures.CreateLevel(map[string]interface{}{
		"NextLevelXpRequirement": 100,
		"LevelNumber":            10001,
	})
	user := userFixtures.CreateUser(map[string]interface{}{
		"Xp":             0,
		"LevelId":        userLevel.Id,
		"CurrencyAmount": 20,
	})

	gameFixtures.CreateUserQuest(user, quest)

	uq := gamificationModels.UserQuestProgressRequest{
		UserUuid: user.Uuid,
		QuestId:  quest.Id,
	}

	fixtures.MockJsonPost(c, uq)

	gamificationController.UpdateUserQuestProgress(c)

	// Assert that user has leveled up
	newUser := userFixtures.GetUserByID(user.Id)
	assert.Equal(t.T(), 200, w.Code)
	assert.Equal(t.T(), nextLevel.Id, newUser.LevelId, "User should have leveled up")
	assert.Equal(t.T(), userLevel.CurrencyWon+user.CurrencyAmount, newUser.CurrencyAmount, "User should get currency from level")
}

func (t *UserProgressTestSuite) TestIShouldProgressQuestWithProgressHandler() {
	c, w = CreateGinTestContext()
	data := map[string]interface{}{
		"DoneCondition": 3,
	}
	quest := gameFixtures.CreateQuest(data)
	user := userFixtures.CreateUser(map[string]interface{}{})
	gameFixtures.CreateUserQuest(user, quest)

	uq := gamificationModels.UserQuestProgressRequest{
		UserUuid: user.Uuid,
		QuestId:  quest.Id,
	}

	fixtures.MockJsonPatch(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUserQuest := gameFixtures.GetUserQuest(quest.Id, user.Id)

	assert.Equal(t.T(), w.Code, 200)
	assert.Equal(t.T(), newUserQuest.Progression, 1)

	c, w = CreateGinTestContext()

	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUserQuest = gameFixtures.GetUserQuest(quest.Id, user.Id)

	assert.Equal(t.T(), w.Code, 200)
	assert.Equal(t.T(), newUserQuest.Progression, 2)

	c, w = CreateGinTestContext()
	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUserQuest = gameFixtures.GetUserQuest(quest.Id, user.Id)

	assert.Equal(t.T(), w.Code, 200, "I should get a 200 code")
	assert.Equal(t.T(), newUserQuest.Progression, 3, "I should get the progression of the user quest to 3")
	assert.Equal(t.T(), newUserQuest.IsCompleted, true)
}

func (t *UserProgressTestSuite) TestQuestShouldNotBeCompletedIfProgressionIsNotMet() {
	c, w = CreateGinTestContext()
	data := map[string]interface{}{
		"DoneCondition": 3,
	}
	quest := gameFixtures.CreateQuest(data)

	user := userFixtures.CreateUser(map[string]interface{}{})
	gameFixtures.CreateUserQuest(user, quest)

	uq := gamificationModels.UserQuestProgressRequest{
		UserUuid: user.Uuid,
		QuestId:  quest.Id,
	}

	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUserQuest := gameFixtures.GetUserQuest(quest.Id, user.Id)

	assert.Equal(t.T(), newUserQuest.Progression, 1)

	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUserQuest = gameFixtures.GetUserQuest(quest.Id, user.Id)

	assert.Equal(t.T(), newUserQuest.Progression, 2)
	if newUserQuest.IsCompleted {
		t.T().Errorf("Quest should not be completed")
	}
	assert.Equal(t.T(), newUserQuest.IsCompleted, false)
}

func (t *UserProgressTestSuite) TestUserShouldGetExperienceFromCompletedQuest() {
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

	uq := gamificationModels.UserQuestProgressRequest{
		UserUuid: user.Uuid,
		QuestId:  quest.Id,
	}

	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUser := userFixtures.GetUserByID(user.Id)

	if newUser.Xp != 0 {
		t.T().Errorf("Quest is not completed and should not give xp")
	}
	assert.Equal(t.T(), newUser.Xp, 0)

	fixtures.MockJsonPost(c, uq)
	gamificationController.UpdateUserQuestProgress(c)

	newUser = userFixtures.GetUserByID(user.Id)

	if newUser.Xp != 5 {
		t.T().Errorf("Quest is completed and should give xp")
	}
	assert.Equal(t.T(), newUser.Xp, 5)
}

func (t *UserProgressTestSuite) TestUserSuccessShouldProgressWhenRelatedQuestIsCompleted() {
	c, w = CreateGinTestContext()

	// Setup everything
	tag := gameFixtures.CreateTag(map[string]interface{}{
		"Name": gamificationModels.PlayQuizTag,
	})
	quest := gameFixtures.CreateQuest(map[string]interface{}{
		"TagId":         tag.Id,
		"Xp":            10,
		"DoneCondition": 1,
	})
	success := gameFixtures.CreateSuccess(map[string]interface{}{
		"TagId":          tag.Id,
		"CurrencyReward": 10,
		"Xp":             10,
	})
	user := userFixtures.CreateUser(map[string]interface{}{
		"Xp": 0,
	})

	// Create user quest and user success
	gameFixtures.CreateUserQuest(user, quest)
	gameFixtures.CreateUserSuccess(user, success)

	uq := gamificationModels.UserQuestProgressRequest{
		UserUuid: user.Uuid,
		QuestId:  quest.Id,
	}

	fixtures.MockJsonPost(c, uq)

	gamificationController.UpdateUserQuestProgress(c)

	progressedUserSuccess := gameFixtures.GetUserSuccess(user.Id, success.Id)

	assert.Equal(t.T(), w.Code, 200)
	assert.Equal(t.T(), progressedUserSuccess.Progression, 1)
}

func (t *UserProgressTestSuite) TestUserShouldGetSuccessXpWhenSuccessIsCompleted() {
	c, w = CreateGinTestContext()

	// Setup everything
	tag := gameFixtures.CreateTag(map[string]interface{}{
		"Name": gamificationModels.PlayQuizTag,
	})
	quest := gameFixtures.CreateQuest(map[string]interface{}{
		"TagId":         tag.Id,
		"Xp":            10,
		"DoneCondition": 1,
	})
	success := gameFixtures.CreateSuccess(map[string]interface{}{
		"TagId":          tag.Id,
		"CurrencyReward": 10,
		"Xp":             10,
		"DoneCondition":  1,
	})
	// Create a level to be sure that user is not max level
	userLevel := gameFixtures.CreateLevel(map[string]interface{}{
		"NextLevelXpRequirement": 100,
		"LevelNumber":            10000,
	})
	user := userFixtures.CreateUser(map[string]interface{}{
		"Xp":      0,
		"LevelId": userLevel.Id,
	})

	// Create user quest and user success
	gameFixtures.CreateUserQuest(user, quest)
	gameFixtures.CreateUserSuccess(user, success)

	uq := gamificationModels.UserQuestProgressRequest{
		UserUuid: user.Uuid,
		QuestId:  quest.Id,
	}

	fixtures.MockJsonPost(c, uq)

	gamificationController.UpdateUserQuestProgress(c)

	progressedUserSuccess := gameFixtures.GetUserSuccess(user.Id, success.Id)
	user = userFixtures.GetUserByID(user.Id)

	assert.Equal(t.T(), 200, w.Code)
	assert.Equal(t.T(), true, progressedUserSuccess.IsCompleted, "User success should be completed")
	assert.Equal(t.T(), user.Xp, quest.Xp+success.Xp, "User should get xp from quest and success")
}

func (t *UserProgressTestSuite) TestUserShouldGetCurrencyFromCompletedSuccess() {
	c, w = CreateGinTestContext()

	// Setup everything
	tag := gameFixtures.CreateTag(map[string]interface{}{
		"Name": gamificationModels.PlayQuizTag,
	})
	quest := gameFixtures.CreateQuest(map[string]interface{}{
		"TagId":         tag.Id,
		"Xp":            10,
		"DoneCondition": 1,
	})
	success := gameFixtures.CreateSuccess(map[string]interface{}{
		"TagId":          tag.Id,
		"CurrencyReward": 10,
		"Xp":             10,
		"DoneCondition":  1,
	})
	// Create a level to be sure that user is not max level
	userLevel := gameFixtures.CreateLevel(map[string]interface{}{
		"NextLevelXpRequirement": 100,
		"LevelNumber":            10000,
	})
	user := userFixtures.CreateUser(map[string]interface{}{
		"Xp":      0,
		"LevelId": userLevel.Id,
	})

	// Create user quest and user success
	gameFixtures.CreateUserQuest(user, quest)
	gameFixtures.CreateUserSuccess(user, success)

	uq := gamificationModels.UserQuestProgressRequest{
		UserUuid: user.Uuid,
		QuestId:  quest.Id,
	}

	fixtures.MockJsonPost(c, uq)

	gamificationController.UpdateUserQuestProgress(c)

	progressedUserSuccess := gameFixtures.GetUserSuccess(user.Id, success.Id)
	user = userFixtures.GetUserByID(user.Id)

	assert.Equal(t.T(), 200, w.Code)
	assert.Equal(t.T(), true, progressedUserSuccess.IsCompleted, "User success should be completed")
	assert.Equal(t.T(), user.CurrencyAmount, success.CurrencyReward, "User should get currency from success")
}

func (t *UserProgressTestSuite) TestUserShouldGetNextRankSuccessWhenUserSuccessIsCompleted() {
	c, w = CreateGinTestContext()

	// Setup everything
	tag := gameFixtures.CreateTag(map[string]interface{}{
		"Name": gamificationModels.PlayQuizTag,
	})
	quest := gameFixtures.CreateQuest(map[string]interface{}{
		"TagId":         tag.Id,
		"Xp":            10,
		"DoneCondition": 1,
	})
	success := gameFixtures.CreateSuccess(map[string]interface{}{
		"TagId":           tag.Id,
		"CurrencyReward":  10,
		"Xp":              10,
		"DoneCondition":   1,
		"ProgressionRank": 1,
	})
	nextSuccess := gameFixtures.CreateSuccess(map[string]interface{}{
		"TagId":           tag.Id,
		"CurrencyReward":  10,
		"Xp":              10,
		"DoneCondition":   1,
		"ProgressionRank": 2,
	})
	// Create a level to be sure that user is not max level
	userLevel := gameFixtures.CreateLevel(map[string]interface{}{
		"NextLevelXpRequirement": 100,
		"LevelNumber":            10000,
	})
	user := userFixtures.CreateUser(map[string]interface{}{
		"Xp":      0,
		"LevelId": userLevel.Id,
	})

	// Create user quest and user success
	gameFixtures.CreateUserQuest(user, quest)
	gameFixtures.CreateUserSuccess(user, success)

	uq := gamificationModels.UserQuestProgressRequest{
		UserUuid: user.Uuid,
		QuestId:  quest.Id,
	}

	fixtures.MockJsonPost(c, uq)

	gamificationController.UpdateUserQuestProgress(c)

	doneSuccess := gameFixtures.GetUserSuccess(success.Id, user.Id)
	user = userFixtures.GetUserByID(user.Id)

	assert.Equal(t.T(), 200, w.Code)
	assert.Equal(t.T(), true, doneSuccess.IsCompleted, "Confirmed user success should be completed")

	// Check if user has next success
	nextUserSuccess := gameFixtures.GetUserSuccess(nextSuccess.Id, user.Id)

	assert.Equal(t.T(), nextSuccess.Id, nextUserSuccess.SuccessId, "User should have next success")
}

func (t *UserProgressTestSuite) TestUserXpShouldBeTheDifferenceOfCurrentXpAndNextLevelXpWhenLevelUpWithASuccess() {
	c, w = CreateGinTestContext()

	tag := gameFixtures.CreateTag(map[string]interface{}{
		"Name": gamificationModels.PlayQuizTag,
	})
	quest := gameFixtures.CreateQuest(map[string]interface{}{
		"TagId":         tag.Id,
		"Xp":            0,
		"DoneCondition": 1,
	})
	success := gameFixtures.CreateSuccess(map[string]interface{}{
		"TagId":          tag.Id,
		"CurrencyReward": 10,
		"Xp":             20,
		"DoneCondition":  1,
	})
	userLevel := gameFixtures.CreateLevel(map[string]interface{}{
		"NextLevelXpRequirement": 10,
		"LevelNumber":            10000,
	})
	nextLevel := gameFixtures.CreateLevel(map[string]interface{}{
		"NextLevelXpRequirement": 100,
		"LevelNumber":            10001,
	})
	user := userFixtures.CreateUser(map[string]interface{}{
		"Xp":      0,
		"LevelId": userLevel.Id,
	})

	// Create user quest and user success
	gameFixtures.CreateUserQuest(user, quest)
	gameFixtures.CreateUserSuccess(user, success)

	uq := gamificationModels.UserQuestProgressRequest{
		UserUuid: user.Uuid,
		QuestId:  quest.Id,
	}

	fixtures.MockJsonPost(c, uq)

	gamificationController.UpdateUserQuestProgress(c)

	user = userFixtures.GetUserByID(user.Id)
	assert.Equal(t.T(), 200, w.Code)
	assert.Equal(t.T(), nextLevel.Id, user.LevelId, "User should have leveled up")
	assert.Equal(t.T(), success.Xp-userLevel.NextLevelXpRequirement, user.Xp, "User should get xp from success")
}
