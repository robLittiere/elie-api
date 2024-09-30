package tests

import (
	"elie-api/modules/fixtures"
	"elie-api/modules/fixtures/common"
	gamificationFixtures "elie-api/modules/fixtures/gamification"
	userFixtures "elie-api/modules/fixtures/user"
	"elie-api/modules/gamification/models"
	successes2 "elie-api/modules/gamification/successes/UpdateUserSuccess"
	models2 "elie-api/modules/gamification/successes/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserSuccessProgressTestSuite struct {
	suite.Suite
}

func (s *UserSuccessProgressTestSuite) SetupTest() {
	Init()
	common.CreateBatchLevels()
}

func TestUserSuccessProgressTestSuite(t *testing.T) {
	suite.Run(t, new(UserSuccessProgressTestSuite))
}

func (s *UserSuccessProgressTestSuite) TestUserShouldProgressSuccessWithSuccessProgressHandler() {
	c, w = CreateGinTestContext()
	userCreated := userFixtures.CreateUser(map[string]interface{}{})
	// Lets test with a login tag
	tag := gamificationFixtures.CreateTag(map[string]interface{}{
		"Name": models.ConnectionTag,
	})
	success := gamificationFixtures.CreateSuccess(map[string]interface{}{
		"Name":            "LoginSuccess",
		"TagId":           tag.Id,
		"DoneProgression": 5,
		"ProgressionRank": 4,
	})

	// Attribute success to user
	gamificationFixtures.CreateUserSuccess(userCreated, success, map[string]interface{}{
		"Progression": 1,
	})

	usr := models2.UserSuccessProgressRequest{
		UserUuid:  userCreated.Uuid,
		SuccessId: success.Id,
	}
	fixtures.MockJsonPatch(c, usr)

	// User progresses success
	successes2.UpdateUserSuccessProgress(c)

	updatedUserSuccess := gamificationFixtures.GetUserSuccess(success.Id, userCreated.Id)
	assert.Equal(s.T(), 2, updatedUserSuccess.Progression, "User should have progressed the success")

	// User progresses success once more
	c, w = CreateGinTestContext()
	fixtures.MockJsonPost(c, usr)
	successes2.UpdateUserSuccessProgress(c)

	// Get user success
	updatedUserSuccess = gamificationFixtures.GetUserSuccess(success.Id, userCreated.Id)
	assert.Equal(s.T(), 3, updatedUserSuccess.Progression, "User should have progressed the success")
}

func (t *UserSuccessProgressTestSuite) TestUserShouldGetCurrencyFromCompletedSuccess() {
	c, w = CreateGinTestContext()

	// Setup everything
	tag := gamificationFixtures.CreateTag(map[string]interface{}{
		"Name": models.PlayGameTag,
	})
	success := gamificationFixtures.CreateSuccess(map[string]interface{}{
		"TagId":          tag.Id,
		"CurrencyReward": 10,
		"Xp":             10,
		"DoneCondition":  1,
	})
	// Create a level to be sure that userCreated is not max level
	userLevel := gamificationFixtures.CreateLevel(map[string]interface{}{
		"NextLevelXpRequirement": 100,
		"LevelNumber":            10000,
	})
	userCreated := userFixtures.CreateUser(map[string]interface{}{
		"Xp":      0,
		"LevelId": userLevel.Id,
	})

	// Create userCreated quest and userCreated success
	gamificationFixtures.CreateUserSuccess(userCreated, success, map[string]interface{}{})

	usr := models2.UserSuccessProgressRequest{
		UserUuid:  userCreated.Uuid,
		SuccessId: success.Id,
	}

	fixtures.MockJsonPatch(c, usr)

	successes2.UpdateUserSuccessProgress(c)

	progressedUserSuccess := gamificationFixtures.GetUserSuccess(userCreated.Id, success.Id)
	userCreated = userFixtures.GetUserByID(userCreated.Id)

	assert.Equal(t.T(), 200, w.Code)
	assert.Equal(t.T(), true, progressedUserSuccess.IsCompleted, "User success should be completed")
	assert.Equal(t.T(), userCreated.CurrencyAmount, success.CurrencyReward, "User should get currency from success")
}

func (t *UserSuccessProgressTestSuite) TestUserShouldGetNextRankSuccessWhenUserSuccessIsCompleted() {
	c, w = CreateGinTestContext()

	// Setup everything
	tag := gamificationFixtures.CreateTag(map[string]interface{}{
		"Name": models.PlayQuizTag,
	})
	success := gamificationFixtures.CreateSuccess(map[string]interface{}{
		"TagId":           tag.Id,
		"CurrencyReward":  10,
		"Xp":              10,
		"DoneCondition":   1,
		"ProgressionRank": 1,
	})
	nextSuccess := gamificationFixtures.CreateSuccess(map[string]interface{}{
		"TagId":           tag.Id,
		"CurrencyReward":  10,
		"Xp":              10,
		"DoneCondition":   1,
		"ProgressionRank": 2,
	})
	// Create a level to be sure that user is not max level
	userLevel := gamificationFixtures.CreateLevel(map[string]interface{}{
		"NextLevelXpRequirement": 100,
		"LevelNumber":            10000,
	})
	user := userFixtures.CreateUser(map[string]interface{}{
		"Xp":      0,
		"LevelId": userLevel.Id,
	})

	// Create user quest and user success
	gamificationFixtures.CreateUserSuccess(user, success, map[string]interface{}{})

	usr := models2.UserSuccessProgressRequest{
		UserUuid:  user.Uuid,
		SuccessId: success.Id,
	}

	fixtures.MockJsonPost(c, usr)

	successes2.UpdateUserSuccessProgress(c)

	doneSuccess := gamificationFixtures.GetUserSuccess(success.Id, user.Id)
	user = userFixtures.GetUserByID(user.Id)

	assert.Equal(t.T(), 200, w.Code)
	assert.Equal(t.T(), true, doneSuccess.IsCompleted, "Confirmed user success should be completed")

	// Check if user has next success
	nextUserSuccess := gamificationFixtures.GetUserSuccess(nextSuccess.Id, user.Id)

	assert.Equal(t.T(), nextSuccess.Id, nextUserSuccess.SuccessId, "User should have next success")
}

func (t *UserSuccessProgressTestSuite) TestUserXpShouldBeTheDifferenceOfCurrentXpAndNextLevelXpWhenLevelUpWithASuccess() {
	c, w = CreateGinTestContext()

	tag := gamificationFixtures.CreateTag(map[string]interface{}{
		"Name": models.PlayQuizTag,
	})
	success := gamificationFixtures.CreateSuccess(map[string]interface{}{
		"TagId":          tag.Id,
		"CurrencyReward": 10,
		"Xp":             20,
		"DoneCondition":  1,
	})
	userLevel := gamificationFixtures.CreateLevel(map[string]interface{}{
		"NextLevelXpRequirement": 10,
		"LevelNumber":            10000,
	})
	nextLevel := gamificationFixtures.CreateLevel(map[string]interface{}{
		"NextLevelXpRequirement": 100,
		"LevelNumber":            10001,
	})
	user := userFixtures.CreateUser(map[string]interface{}{
		"Xp":      0,
		"LevelId": userLevel.Id,
	})

	// Create user success
	gamificationFixtures.CreateUserSuccess(user, success, map[string]interface{}{})

	usr := models2.UserSuccessProgressRequest{
		UserUuid:  user.Uuid,
		SuccessId: success.Id,
	}

	fixtures.MockJsonPost(c, usr)

	successes2.UpdateUserSuccessProgress(c)

	user = userFixtures.GetUserByID(user.Id)
	assert.Equal(t.T(), 200, w.Code)
	assert.Equal(t.T(), nextLevel.Id, user.LevelId, "User should have leveled up")
	assert.Equal(t.T(), success.Xp-userLevel.NextLevelXpRequirement, user.Xp, "User should get xp from success")
}
