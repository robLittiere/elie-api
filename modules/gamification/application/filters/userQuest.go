package filters

import (
	"elie-api/modules/common/filter"
	"elie-api/modules/gamification/domain/query"
)

func GetUserQuestFilterRegitry() filter.Filters {
	return filter.Filters{
		{
			Name:          "user_id",
			Criteria:      &query.UserQuestUserIdCriteria{Field: "user_id"},
			Documentation: "Filter by user id",
		},
		{
			Name:          "user_uuid",
			Criteria:      &query.UserQuestWithUserUuid{Field: "user_uuid"},
			Documentation: "Filter by user uuid",
		},
	}

}
