package filters

import (
	"elie-api/modules/common/filter"
	"elie-api/modules/game/domain/query"
)

func GetQuizFilterRegistry() filter.Filters {
	return filter.Filters{
		{
			Name:          "topic",
			Criteria:      &query.QuizTopicCriteria{Field: "topic"},
			Documentation: "Filter by quiz topic",
		},
	}
}
