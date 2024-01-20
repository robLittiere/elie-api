package filters

import (
	"elie-api/modules/common/filter"
	"elie-api/modules/game/domain/query"
)

func GetQuizFilterRegistry() filter.Filters {
	return filter.Filters{
		{
			Name:          "topic_name",
			Criteria:      &query.QuizTopicNameCriteria{Field: "topic"},
			Documentation: "Filter by quiz topic",
		},
		{
			Name:          "topic_id",
			Criteria:      &query.QuizTopicIdCriteria{Field: "topic"},
			Documentation: "Filter by quiz topic id",
		},
		{
			Name:          "topic_id[]",
			Criteria:      &query.QuizTopicIdArrayCriteria{Field: "topic"},
			Documentation: "Filter by quiz topic id array",
		},
		{
			Name:          "topic_name[]",
			Criteria:      &query.QuizTopicNameArrayCriteria{Field: "topic"},
			Documentation: "Filter by quiz topic name array",
		},
		{
			Name:          "quiz_id",
			Criteria:      &query.QuizIdCriteria{Field: "quiz"},
			Documentation: "Filter by quiz id",
		},
		{
			Name:          "title",
			Criteria:      &query.QuizTitleCriteria{Field: "quiz"},
			Documentation: "Filter by quiz title",
		},
	}
}
