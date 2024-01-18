package filters

import (
	"elie-api/modules/common/filter"
	"elie-api/modules/game/domain/query"
)

func GetQuizFilterRegistry() filter.Filters {
	return filter.Filters{
		{
			Name:          "topicName",
			Criteria:      &query.QuizTopicNameCriteria{Field: "topic"},
			Documentation: "Filter by quiz topic",
		},
		{
			Name:          "topicId",
			Criteria:      &query.QuizTopicIdCriteria{Field: "topic"},
			Documentation: "Filter by quiz topic id",
		},
		{
			Name:          "topicId[]",
			Criteria:      &query.QuizTopicIdArrayCriteria{Field: "topic"},
			Documentation: "Filter by quiz topic id array",
		},
		{
			Name:          "topicName[]",
			Criteria:      &query.QuizTopicNameArrayCriteria{Field: "topic"},
			Documentation: "Filter by quiz topic name array",
		},
	}
}
