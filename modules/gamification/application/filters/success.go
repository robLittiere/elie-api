package filters

import (
	"elie-api/modules/common/filter"
	"elie-api/modules/gamification/domain/query"
)

func GetSuccessFilterRegitry() filter.Filters {
	return filter.Filters{
		{
			Name:          "tag_id",
			Criteria:      &query.SuccessWithTagIdCriteria{Field: "tag_id"},
			Documentation: "Get successes by their tag id",
		},
		{
			Name:          "tag_name",
			Criteria:      &query.SuccessWithTagNameCriteria{Field: "tag_name"},
			Documentation: "Get successes by their tag name",
		},
	}
}
