package filters

import (
	"elie-api/modules/common/filter"
	query2 "elie-api/modules/user/domain/query"
)

// GetUserFilters Get the list of available filters for a given model
func GetUserFilters() filter.Filters {
	return filter.Filters{
		{
			Name:          "username",
			Criteria:      &query2.UsernameCriteria{Field: "username"},
			Documentation: "Filter by username",
		},
		{
			Name:          "email",
			Criteria:      &query2.EmailCriteria{Field: "email"},
			Documentation: "Filter by email",
		},
		{
			Name:          "uid",
			Criteria:      &query2.UuidCriteria{Field: "uid"},
			Documentation: "Filter by uid",
		},
	}
}
