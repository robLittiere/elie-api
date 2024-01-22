package tests

import (
	"elie-api/modules/common/criteria"
	"elie-api/modules/common/repository"
	"elie-api/modules/user/application/filters"
	"elie-api/modules/user/domain/query"
	"reflect"
	"testing"
)

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
