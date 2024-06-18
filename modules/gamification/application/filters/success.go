package filters

import (
	"elie-api/modules/common/filter"
	"elie-api/modules/gamification/domain/query"
)

func GetSuccessFilterRegitry() filter.Filters {
	return filter.Filters{
		{
			Name:          "tag",
			Criteria:      &query.SuccessTagCriteria{Field: "tag"},
			Documentation: "Get successes by their tag name",
		},
	}
}
