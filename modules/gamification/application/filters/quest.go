package filters

import (
	"elie-api/modules/common/filter"
	"elie-api/modules/gamification/domain/query"
)

func GetQuestFilterRegistry() filter.Filters {
	return filter.Filters{
		{
			Name:          "tag",
			Criteria:      &query.QuestTagCriteria{Field: "tag"},
			Documentation: "Get quests by their tag name",
		},
	}
}
