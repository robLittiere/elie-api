package tests

import (
	"elie-api/modules/fixtures"
	"elie-api/modules/fixtures/common"
	gamificationFixtures "elie-api/modules/fixtures/gamification"
	userFixtures "elie-api/modules/fixtures/user"
	"elie-api/modules/gamification/successes/domain/models"
	"elie-api/modules/gamification/successes/listUserSuccesses"
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
	listUserSuccesses.GetUserSuccesses(c)

	// Assert we should only get the success with the tag we asked for
	var suc []models.Success
	err := json.NewDecoder(w.Body).Decode(&suc)
	if err != nil {
		s.T().Errorf("Error while decoding response : %v", err)
	}

	assert.Equal(s.T(), 200, w.Code, "I should get a 200 code")
	assert.Equal(s.T(), 1, len(suc), "I should get one success")
	assert.Equal(s.T(), success.Id, suc[0].Id, "I should get the successes of the user")
}

func (s *UserSuccessTestSuite) TestIShouldGetEmptyListIfUserHasNoSuccesses() {
	c, w = CreateGinTestContext()
	userCreated := userFixtures.CreateUser(map[string]interface{}{})

	u := url.Values{}
	u.Add("user_uuid", userCreated.Uuid)

	fixtures.MockJsonGet(c, nil, u)
	listUserSuccesses.GetUserSuccesses(c)

	// Assert we should only get the success with the tag we asked for
	var suc []models.Success
	err := json.NewDecoder(w.Body).Decode(&suc)
	if err != nil {
		s.T().Errorf("Error while decoding response : %v", err)
	}

	assert.Equal(s.T(), 200, w.Code, "I should get a 200 code")
	assert.Equal(s.T(), 0, len(suc), "I should get no success")
}
