package filters

import (
	"elie-api/modules/common/filter"
	"elie-api/modules/gamification/domain/query"
)

func GetQuestFilterRegistry() filter.Filters {
	return filter.Filters{
		{
			Name:          "type",
			Criteria:      &query.QuestTypeNameCriteria{Field: "type"},
			Documentation: "Get quests by their type name",
		},
	}
}
