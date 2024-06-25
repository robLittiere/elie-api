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

type UserSuccessTestSuite struct {
	suite.Suite
}

func (s *UserSuccessTestSuite) SetupTest() {
	Init()
}

func TestUserSuccessTestSuite(t *testing.T) {
	suite.Run(t, new(UserSuccessTestSuite))
}

func (s *UserSuccessTestSuite) TestIShouldGetSuccessesWithAUserUuid() {
	c, w = CreateGinTestContext()
	userCreated := user.CreateUser(map[string]interface{}{})
	success := gamificationFixtures.CreateSuccess(map[string]interface{}{
		"Name": "test success",
	})

	gamificationFixtures.CreateUserSuccess(userCreated, success)

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
	userCreated := user.CreateUser(map[string]interface{}{})

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
