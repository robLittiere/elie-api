package filters

import (
	"elie-api/modules/common/filter"
	"elie-api/modules/gamification/domain/query"
)

func GetUserSuccessFilterRegitry() filter.Filters {
	return filter.Filters{
		{
			Name:          "user_id",
			Criteria:      &query.UserSuccessUserIdCriteria{Field: "user_id"},
			Documentation: "Filter by user id",
		},
	}

}
