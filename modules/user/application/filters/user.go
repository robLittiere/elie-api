package filters

import (
	"elie-api/modules/common/criteria"
	"elie-api/modules/user/domain/query"
	"fmt"
)

// Filters are a list of the queryParam name, the criteria associated to it and a documentation for what is the param
type Filters []Filter

// Filter is a struct that contains the queryParam name, the criteria associated to it and a documentation for what is the param
type Filter struct {
	Name          string
	Criteria      criteria.Criteria
	Documentation string
}

// GetUserFilters Get the list of available filters for a given model
func GetUserFilters() Filters {
	return Filters{
		{
			Name:          "username",
			Criteria:      &query.UsernameCriteria{Field: "username"},
			Documentation: "Filter by username",
		},
		{
			Name:          "email",
			Criteria:      &query.EmailCriteria{Field: "email"},
			Documentation: "Filter by email",
		},
		{
			Name:          "uid",
			Criteria:      &query.UuidCriteria{Field: "uid"},
			Documentation: "Filter by uid",
		},
	}
}

func GetUserFilter(s string) (criteria.Criteria, error) {
	filters := GetUserFilters()

	// get the filter according to its name
	for _, filter := range filters {
		if filter.Name == s {
			return filter.Criteria, nil
		}
	}
	return nil, fmt.Errorf("query parameter %v is not a valid filter", s)
}
