package tests

import (
	"elie-api/modules/fixtures"
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
}

func TestUserQuestTestSuite(t *testing.T) {
	suite.Run(t, new(UserQuestTestSuite))
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
	var quests []models.Quest
	err := json.NewDecoder(w.Body).Decode(&quests)
	if err != nil {
		t.T().Errorf("Error while decoding response : %v", err)
	}

	assert.Equal(t.T(), 200, w.Code, "I should get a 200 code")
	assert.Equal(t.T(), 1, len(quests), "I should get one quest")
	assert.Equal(t.T(), quest.Id, quests[0].Id, "I should get the quests of the user")
}

func (t *UserQuestTestSuite) TestIShouldGetEmptyArrayIfUserHasNoQuests() {
	c, w = CreateGinTestContext()
	userCreated := user.CreateUser(map[string]interface{}{})

	u := url.Values{}
	u.Add("user_uuid", userCreated.Uuid)

	fixtures.MockJsonGet(c, nil, u)
	gamificationController.GetUserQuests(c)

	// Assert we should only get the quest with the tag we asked for
	var quests []models.Quest
	err := json.NewDecoder(w.Body).Decode(&quests)
	if err != nil {
		t.T().Errorf("Error while decoding response : %v", err)
	}

	assert.Equal(t.T(), 200, w.Code, "I should get a 200 code")
	assert.Equal(t.T(), 0, len(quests), "I should get an empty array")
}
