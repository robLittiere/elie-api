package tests

import (
	"elie-api/modules/fixtures"
	"elie-api/modules/fixtures/common"
	gamificationFixtures "elie-api/modules/fixtures/gamification"
	userFixtures "elie-api/modules/fixtures/user"
	gamificationController "elie-api/modules/gamification/application/controllers"
	"elie-api/modules/gamification/models"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/url"
	"testing"
)

type UserSuccessTestSuite struct {
	suite.Suite
}

func (s *UserSuccessTestSuite) SetupTest() {
	Init()
	common.CreateBatchLevels()
}

func TestUserSuccessTestSuite(t *testing.T) {
	suite.Run(t, new(UserSuccessTestSuite))
}

func (s *UserSuccessTestSuite) TestIShouldGetSuccessesWithAUserUuid() {
	c, w = CreateGinTestContext()
	userCreated := userFixtures.CreateUser(map[string]interface{}{})
	success := gamificationFixtures.CreateSuccess(map[string]interface{}{
		"Name": "test success",
	})

	gamificationFixtures.CreateUserSuccess(userCreated, success, map[string]interface{}{})

	u := url.Values{}
	u.Add("user_uuid", userCreated.Uuid)

	fixtures.MockJsonGet(c, nil, u)
	gamificationController.GetUserSuccesses(c)

	// Assert we should only get the success with the tag we asked for
	var successes []models.Success
	err := json.NewDecoder(w.Body).Decode(&successes)
	if err != nil {
		s.T().Errorf("Error while decoding response : %v", err)
	}

	assert.Equal(s.T(), 200, w.Code, "I should get a 200 code")
	assert.Equal(s.T(), 1, len(successes), "I should get one success")
	assert.Equal(s.T(), success.Id, successes[0].Id, "I should get the successes of the user")
}

func (s *UserSuccessTestSuite) TestIShouldGetEmptyListIfUserHasNoSuccesses() {
	c, w = CreateGinTestContext()
	userCreated := userFixtures.CreateUser(map[string]interface{}{})

	u := url.Values{}
	u.Add("user_uuid", userCreated.Uuid)

	fixtures.MockJsonGet(c, nil, u)
	gamificationController.GetUserSuccesses(c)

	// Assert we should only get the success with the tag we asked for
	var successes []models.Success
	err := json.NewDecoder(w.Body).Decode(&successes)
	if err != nil {
		s.T().Errorf("Error while decoding response : %v", err)
	}

	assert.Equal(s.T(), 200, w.Code, "I should get a 200 code")
	assert.Equal(s.T(), 0, len(successes), "I should get no success")
}

func (s *UserSuccessTestSuite) TestUserShouldProgressSuccessWithSuccessProgressHandler() {
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

	usr := models.UserSuccessProgressRequest{
		UserUuid:  userCreated.Uuid,
		SuccessId: success.Id,
	}
	fixtures.MockJsonPatch(c, usr)

	// User progresses success
	gamificationController.UpdateUserSuccessProgress(c)

	updatedUserSuccess := gamificationFixtures.GetUserSuccess(success.Id, userCreated.Id)
	assert.Equal(s.T(), 2, updatedUserSuccess.Progression, "User should have progressed the success")

	// User progresses success once more
	c, w = CreateGinTestContext()
	fixtures.MockJsonPost(c, usr)
	gamificationController.UpdateUserSuccessProgress(c)

	// Get user success
	updatedUserSuccess = gamificationFixtures.GetUserSuccess(success.Id, userCreated.Id)
	assert.Equal(s.T(), 3, updatedUserSuccess.Progression, "User should have progressed the success")
}

func (t *UserSuccessTestSuite) TestUserShouldGetCurrencyFromCompletedSuccess() {
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

	usr := models.UserSuccessProgressRequest{
		UserUuid:  userCreated.Uuid,
		SuccessId: success.Id,
	}

	fixtures.MockJsonPatch(c, usr)

	gamificationController.UpdateUserSuccessProgress(c)

	progressedUserSuccess := gamificationFixtures.GetUserSuccess(userCreated.Id, success.Id)
	userCreated = userFixtures.GetUserByID(userCreated.Id)

	assert.Equal(t.T(), 200, w.Code)
	assert.Equal(t.T(), true, progressedUserSuccess.IsCompleted, "User success should be completed")
	assert.Equal(t.T(), userCreated.CurrencyAmount, success.CurrencyReward, "User should get currency from success")
}

func (t *UserProgressTestSuite) TestUserShouldGetNextRankSuccessWhenUserSuccessIsCompleted() {
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

	usr := models.UserSuccessProgressRequest{
		UserUuid:  user.Uuid,
		SuccessId: success.Id,
	}

	fixtures.MockJsonPost(c, usr)

	gamificationController.UpdateUserSuccessProgress(c)

	doneSuccess := gamificationFixtures.GetUserSuccess(success.Id, user.Id)
	user = userFixtures.GetUserByID(user.Id)

	assert.Equal(t.T(), 200, w.Code)
	assert.Equal(t.T(), true, doneSuccess.IsCompleted, "Confirmed user success should be completed")

	// Check if user has next success
	nextUserSuccess := gamificationFixtures.GetUserSuccess(nextSuccess.Id, user.Id)

	assert.Equal(t.T(), nextSuccess.Id, nextUserSuccess.SuccessId, "User should have next success")
}

func (t *UserProgressTestSuite) TestUserXpShouldBeTheDifferenceOfCurrentXpAndNextLevelXpWhenLevelUpWithASuccess() {
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

	usr := models.UserSuccessProgressRequest{
		UserUuid:  user.Uuid,
		SuccessId: success.Id,
	}

	fixtures.MockJsonPost(c, usr)

	gamificationController.UpdateUserSuccessProgress(c)

	user = userFixtures.GetUserByID(user.Id)
	assert.Equal(t.T(), 200, w.Code)
	assert.Equal(t.T(), nextLevel.Id, user.LevelId, "User should have leveled up")
	assert.Equal(t.T(), success.Xp-userLevel.NextLevelXpRequirement, user.Xp, "User should get xp from success")
}
