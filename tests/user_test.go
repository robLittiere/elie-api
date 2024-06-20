package tests

import (
	"bytes"
	"elie-api/config"
	"elie-api/modules/auth/controllers"
	"elie-api/modules/common/criteria"
	"elie-api/modules/common/repository"
	"elie-api/modules/fixtures"
	userFixtures "elie-api/modules/fixtures/user"
	userController "elie-api/modules/user/application/controllers"
	"elie-api/modules/user/application/filters"
	"elie-api/modules/user/domain/query"
	"elie-api/modules/user/infrastructure"
	"elie-api/modules/user/models"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestIShouldCreateWithSignupUserHandler(t *testing.T) {
	c, w = CreateGinTestContext()
	uq := models.UserRequest{
		Email:    "anothermailnottakenforsure@mail.com",
		Password: "robinoux",
	}

	fixtures.MockJsonPost(c, uq)
	controllers.SignupHandler(c)

	assert.Equal(t, w.Code, http.StatusCreated)

	var user models.User
	userRepo := infrastructure.NewUserRepo(config.DB)
	record := userRepo.DB.Where("email = ?", uq.Email).First(&user)
	if record.Error != nil {
		t.Errorf("Exepected user to be created but found none : %v", record.Error)
	}
}

func TestIShouldCreateUserE2E(t *testing.T) {
	c, w = CreateGinTestContext()
	router = InitRouter()
	uq := models.UserRequest{
		Email:    "mailnottakenforsure@mail.com",
		Password: "robinator",
	}
	jsonv, _ := json.Marshal(uq)

	// This is a e2e test
	req, _ := http.NewRequest("POST", "/api/v1/signup", bytes.NewReader(jsonv))
	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	router.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusCreated)
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code : %d, but got : %d \n", http.StatusCreated, w.Code)
	}

	var user models.User
	userRepo := infrastructure.NewUserRepo(config.DB)
	record := userRepo.DB.Where("email = ?", uq.Email).First(&user)
	if record.Error != nil {
		t.Errorf("Exepected user to be created but found none : %v", record.Error)
	}

	assert.Equal(t, user.Email, uq.Email)

}

func TestIShouldGetUserWithLoginHandler(t *testing.T) {
	// Create a user using signup handler
	uq := models.UserRequest{
		Email:    "rob@mail.com",
		Password: "rob",
	}

	fixtures.MockJsonPost(c, uq)
	controllers.SignupHandler(c)

	// Recreate context to simulate a new request
	c, w = CreateGinTestContext()

	// Login the user
	uq = models.UserRequest{
		Email:    "rob@mail.com",
		Password: "robinoux",
	}
	fixtures.MockJsonPost(c, uq)
	controllers.LoginHandler(c)

	assert.Equal(t, w.Code, http.StatusOK)

	var user models.User
	err := json.NewDecoder(w.Body).Decode(&user)
	if err != nil {
		t.Errorf("Error while decoding response : %v", err)
	}
	assert.Equal(t, user.Email, uq.Email)
}

func TestIShouldGetUserCriterias(t *testing.T) {

	// Declare some fake query parameters
	queryParams := map[string]string{
		"uuid":          "1",
		"username":      "robinou",
		"nonvalidParam": "test",
	}
	usernameCriteria := query.UsernameCriteria{Field: "username"}
	uuidCriteria := query.UuidCriteria{Field: "uuid"}

	expectedCriterias := []criteria.Criteria{&uuidCriteria, &usernameCriteria}
	criteriaList := make([]criteria.Criteria, 0)

	// That should return an error because nonvalidParam is not a valid filter
	userFilterRegistry := filters.GetUserFilters()
	for k, _ := range queryParams {
		filterCriteria, err := repository.GetCriteria(k, userFilterRegistry)
		if err == nil {
			criteriaList = append(criteriaList, filterCriteria)
		}
	}

	// Assert that the criteria list is the same as the expected one
	if !reflect.DeepEqual(criteriaList, expectedCriterias) {
		t.Errorf("Expected criteria %v but got %v", expectedCriterias, criteriaList)
	}

}

func TestIShouldGetAnErrorForANonValidFilter(t *testing.T) {
	userFilterRegistry := filters.GetUserFilters()
	queryParams := map[string]string{
		"NonExistantFitler": "NonExistantValue",
	}

	for k, _ := range queryParams {
		_, err := repository.GetCriteria(k, userFilterRegistry)
		if err == nil {
			t.Errorf("Expected an error but got nil")
		}
	}
}

func TestIShouldGetUserByUsername(t *testing.T) {
	Init()
	c, w = CreateGinTestContext()

	user := userFixtures.CreateUser(map[string]interface{}{"Username": "elrobinator"})

	u := url.Values{}
	u.Add("username", user.Username)
	fixtures.MockJsonGet(c, nil, u)

	userController.GetUsers(c)

	assert.Equal(t, w.Code, http.StatusOK)

	var users []models.User
	err := json.NewDecoder(w.Body).Decode(&users)
	if err != nil {
		t.Errorf("Error while decoding response : %v", err)
	}
	assert.Equal(t, users[0].Username, user.Username)
}
