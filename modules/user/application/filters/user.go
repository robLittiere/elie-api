package filters

import (
	"elie-api/modules/common/filter"
	"elie-api/modules/user/domain/query"
)

// GetUserFilters Get the list of available filters for a given model
func GetUserFilters() filter.Filters {
	return filter.Filters{
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
			Name:          "uuid",
			Criteria:      &query.UuidCriteria{Field: "uuid"},
			Documentation: "Filter by uid",
		},
	}
}
