package tests

import (
	"elie-api/modules/fixtures"
	"elie-api/modules/fixtures/common"
	gamificationFixtures "elie-api/modules/fixtures/gamification"
	"elie-api/modules/fixtures/user"
	gamificationController "elie-api/modules/gamification/application/controllers"
	"elie-api/modules/gamification/models"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/url"
	"testing"
)

type UserQuestTestSuite struct {
	suite.Suite
}

func (s *UserQuestTestSuite) SetupTest() {
	Init()
	common.CreateBatchLevels()
}

func TestUserQuestTestSuite(t *testing.T) {
	suite.Run(t, new(UserQuestTestSuite))
}

func (t *UserQuestTestSuite) TestIShouldCreateUserQuest() {
	c, w = CreateGinTestContext()
	userCreated := user.CreateUser(map[string]interface{}{})
	quest := gamificationFixtures.CreateQuest(map[string]interface{}{})

	request := models.UserQuestProgressRequest{
		UserUuid: userCreated.Uuid,
		QuestId:  quest.Id,
	}

	fixtures.MockJsonPost(c, request)
	gamificationController.CreateUserQuest(c)

	assert.Equal(t.T(), 200, w.Code, "I should get a 200 code")

	// Assert the user quest has been created
	userQuest := gamificationFixtures.GetUserQuest(quest.Id, userCreated.Id)
	assert.Equal(t.T(), userQuest.QuestId, quest.Id, "The user quest should have the right quest id")
	assert.Equal(t.T(), userQuest.UserId, userCreated.Id, "The user quest should have the right user id")

}

func (t *UserQuestTestSuite) TestIShouldGetQuestsWithAUserUuid() {

	c, w = CreateGinTestContext()
	userCreated := user.CreateUser(map[string]interface{}{})
	quest := gamificationFixtures.CreateQuest(map[string]interface{}{
		"Name": "test quest",
	})

	gamificationFixtures.CreateUserQuest(userCreated, quest)

	u := url.Values{}
	u.Add("user_uuid", userCreated.Uuid)

	fixtures.MockJsonGet(c, nil, u)
	gamificationController.GetUserQuests(c)

	// Assert we should only get the quest with the tag we asked for
	var userQuests []models.UserQuest
	err := json.NewDecoder(w.Body).Decode(&userQuests)
	if err != nil {
		t.T().Errorf("Error while decoding response : %v", err)
	}

	assert.Equal(t.T(), 200, w.Code, "I should get a 200 code")
	assert.Equal(t.T(), 1, len(userQuests), "I should get one quest")
	assert.Equal(t.T(), quest.Id, userQuests[0].Id, "I should get the userQuests of the user")
}

func (t *UserQuestTestSuite) TestIShouldGetEmptyArrayIfUserHasNoQuests() {
	c, w = CreateGinTestContext()
	userCreated := user.CreateUser(map[string]interface{}{})

	u := url.Values{}
	u.Add("user_uuid", userCreated.Uuid)

	fixtures.MockJsonGet(c, nil, u)
	gamificationController.GetUserQuests(c)

	// Assert we should only get the quest with the tag we asked for
	var userQuests []models.Quest
	err := json.NewDecoder(w.Body).Decode(&userQuests)
	if err != nil {
		t.T().Errorf("Error while decoding response : %v", err)
	}

	assert.Equal(t.T(), 200, w.Code, "I should get a 200 code")
	assert.Equal(t.T(), 0, len(userQuests), "I should get an empty array")
}
