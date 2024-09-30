package tests

import (
	"elie-api/modules/fixtures"
	gamificationFixtures "elie-api/modules/fixtures/gamification"
	"elie-api/modules/gamification/models"
	models2 "elie-api/modules/gamification/successes/domain/models"
	"elie-api/modules/gamification/successes/listSuccesses"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/url"
	"strconv"
	"testing"
)

type SuccessTestSuite struct {
	suite.Suite
}

func (s *SuccessTestSuite) SetupTest() {
	Init()
}

func TestSuccessTestSuite(t *testing.T) {
	suite.Run(t, new(SuccessTestSuite))
}

func (s *SuccessTestSuite) TestIShouldGetSuccessWithCorrectTag() {

	// Setup 2 tags that we give to 2 different successes
	c, w = CreateGinTestContext()
	levelTag := gamificationFixtures.CreateTag(map[string]interface{}{
		"Name": models.LevelTag,
	})
	playQuizTag := gamificationFixtures.CreateTag(map[string]interface{}{
		"Name": models.PlayQuizTag,
	})

	success := gamificationFixtures.CreateSuccess(map[string]interface{}{
		"TagId": levelTag.Id,
	})
	gamificationFixtures.CreateSuccess(map[string]interface{}{
		"TagId": playQuizTag.Id,
	})

	u := url.Values{}
	u.Add("tag_id", strconv.Itoa(levelTag.Id))

	// Mock API call
	fixtures.MockJsonGet(c, nil, u)
	listSuccesses.GetSuccesses(c)

	// Assert we should only get the success with the tag we asked for
	var successes []models2.Success
	err := json.NewDecoder(w.Body).Decode(&successes)
	if err != nil {
		s.T().Errorf("Error while decoding response : %v", err)
	}

	assert.Equal(s.T(), len(successes), 1)
	assert.Equal(s.T(), successes[0].TagId, success.Id)
	assert.Equal(s.T(), successes[0].Tag.Name, models.LevelTag)
}

func (s *SuccessTestSuite) TestIShouldGetSuccessWithCorrectTagName() {

	// Setup
	c, w = CreateGinTestContext()
	playQuizTag := gamificationFixtures.CreateTag(map[string]interface{}{
		"Name": models.PlayQuizTag,
	})
	levelTag := gamificationFixtures.CreateTag(map[string]interface{}{
		"Name": models.LevelTag,
	})

	success := gamificationFixtures.CreateSuccess(map[string]interface{}{
		"TagId": playQuizTag.Id,
		"Name":  "test success",
	})
	gamificationFixtures.CreateSuccess(map[string]interface{}{
		"TagId": levelTag.Id,
		"Name":  "test success 2",
	})

	u := url.Values{}
	u.Add("tag_name", string(models.PlayQuizTag))

	// Mock API call
	fixtures.MockJsonGet(c, nil, u)
	listSuccesses.GetSuccesses(c)

	// Assert we should only get the success with the tag we asked for
	var successes []models2.Success
	err := json.NewDecoder(w.Body).Decode(&successes)
	if err != nil {
		s.T().Errorf("Error while decoding response : %v", err)
	}
	assert.Equal(s.T(), 1, len(successes), "I should get only one success")
	assert.Equal(s.T(), success.Id, successes[0].Tag.Id, "I should get the success with the correct tag")
	assert.Equal(s.T(), models.PlayQuizTag, successes[0].Tag.Name)
}

func (s *SuccessTestSuite) TestIShouldGetNoSuccessIfSuccessDoesntHaveTag() {

	// Setup
	c, w = CreateGinTestContext()
	playQuizTag := gamificationFixtures.CreateTag(map[string]interface{}{
		"Name": models.PlayQuizTag,
	})

	gamificationFixtures.CreateSuccess(map[string]interface{}{
		"TagId": playQuizTag.Id,
		"Name":  "test success",
	})

	u := url.Values{}
	u.Add("tag_name", string(models.LevelTag))

	// Mock API call
	fixtures.MockJsonGet(c, nil, u)
	listSuccesses.GetSuccesses(c)

	// Assert we should only get the success with the tag we asked for
	var suc []models2.Success
	err := json.NewDecoder(w.Body).Decode(&suc)
	if err != nil {
		s.T().Errorf("Error while decoding response : %v", err)
	}
	assert.Equal(s.T(), 0, len(suc), "I should get no success")
}
