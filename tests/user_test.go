package tests

import (
	"elie-api/modules/common/criteria"
	"elie-api/modules/user/application/filters"
	"elie-api/modules/user/domain/query"
	"reflect"
	"testing"
)

func TestIShouldGetUserCriterias(t *testing.T) {

	// Declare some fake query parameters
	queryParams := map[string]string{
		"uid":           "1",
		"username":      "test",
		"nonvalidParam": "test",
	}
	usernameCriteria := query.UsernameCriteria{Field: "username"}
	uidCriteria := query.UuidCriteria{Field: "uid"}

	expectedCriterias := []criteria.Criteria{&uidCriteria, &usernameCriteria}
	criteriaList := make([]criteria.Criteria, 0)

	// That should return an error because nonvalidParam is not a valid filter
	for k, _ := range queryParams {
		filterCriteria, _ := filters.GetUserFilter(k)
		criteriaList = append(criteriaList, filterCriteria)
	}

	// Assert that the criteria list is the same as the expected one
	if !reflect.DeepEqual(criteriaList, expectedCriterias) {
		t.Errorf("Expected criteria %v but got %v", expectedCriterias, criteriaList)
	}

}

func TestIShouldGetAnErrorForNonValidFilter(t *testing.T) {
	_, err := filters.GetUserFilter("nonvalidParam")
	if err == nil {
		t.Errorf("Expected an error but got nil")
	}
}
