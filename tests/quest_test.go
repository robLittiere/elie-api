package tests

import (
	"elie-api/modules/fixtures"
	gamificationFixtures "elie-api/modules/fixtures/gamification"
	gamificationController "elie-api/modules/gamification/application/controllers"
	"elie-api/modules/gamification/models"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/url"
	"strconv"
	"testing"
)

type QuestTestSuite struct {
	suite.Suite
}

func (t *QuestTestSuite) SetupTest() {
	Init()
	c, w = CreateGinTestContext()
}

func TestQuestTestSuite(t *testing.T) {
	suite.Run(t, new(QuestTestSuite))
}

func (t *QuestTestSuite) TestIShouldGetQuests() {

	// Create 2 quests
	gamificationFixtures.CreateQuest(map[string]interface{}{})
	gamificationFixtures.CreateQuest(map[string]interface{}{})

	// Mock API call
	fixtures.MockJsonGet(c, nil, nil)
	gamificationController.GetQuests(c)

	// Assert we should get all the quests
	var quests []models.Quest
	err := json.NewDecoder(w.Body).Decode(&quests)
	if err != nil {
		t.T().Errorf("Error while decoding response : %v", err)
	}

	assert.Equal(t.T(), len(quests), 2, "I should get all the quests I just created")
}

func (t *QuestTestSuite) TestIShouldGetQuestsWithCorrectTagId() {
	// Setup 1 tag that we give to 1 quest
	tag := gamificationFixtures.CreateTag(map[string]interface{}{
		"Name": models.WonQuizTag,
	})
	quest := gamificationFixtures.CreateQuest(map[string]interface{}{
		"TagId": tag.Id,
	})

	// Mock API call
	u := url.Values{}
	u.Add("tag_id", strconv.Itoa(tag.Id))
	fixtures.MockJsonGet(c, nil, u)
	gamificationController.GetQuests(c)

	// Assert we should only get the quest with the tag we asked for
	var quests []models.Quest
	err := json.NewDecoder(w.Body).Decode(&quests)
	if err != nil {
		t.T().Errorf("Error while decoding response : %v", err)
	}

	assert.Equal(t.T(), len(quests), 1, "I should get only the quest with the tag I asked for")
	assert.Equal(t.T(), quests[0].TagId, quest.Id, "I should get the quest with the tag I asked for")
	assert.Equal(t.T(), quests[0].Tag.Name, models.WonQuizTag, "I should get the quest with the tag I asked for")
}

func (t *QuestTestSuite) TestIShouldGetQuestsWithCorrectTagName() {
	// Setup 2 tag that we give to 2 quest
	tag := gamificationFixtures.CreateTag(map[string]interface{}{
		"Name": models.WonQuizTag,
	})
	uselessTag := gamificationFixtures.CreateTag(map[string]interface{}{
		"Name": models.PlayGameTag,
	})
	quest := gamificationFixtures.CreateQuest(map[string]interface{}{
		"TagId": tag.Id,
	})
	gamificationFixtures.CreateQuest(map[string]interface{}{
		"TagId": uselessTag.Id,
	})

	// Mock API call
	u := url.Values{}
	u.Add("tag_name", string(models.WonQuizTag))
	fixtures.MockJsonGet(c, nil, u)
	gamificationController.GetQuests(c)

	// Assert we should only get the quest with the tag we asked for
	var quests []models.Quest
	err := json.NewDecoder(w.Body).Decode(&quests)
	if err != nil {
		t.T().Errorf("Error while decoding response : %v", err)
	}
	assert.Equal(t.T(), 1, len(quests), "I should get only one quest")
	assert.Equal(t.T(), quests[0].TagId, quest.Id, "I should get the quest with the tag I asked for")
	assert.Equal(t.T(), quests[0].Tag.Name, models.WonQuizTag, "I should get the quest with the tag I asked for")
}
