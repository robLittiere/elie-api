package filters

import (
	"elie-api/modules/common/filter"
	"elie-api/modules/gamification/domain/query"
)

func GetQuestFilterRegistry() filter.Filters {
	return filter.Filters{
		{
			Name:          "tag_id",
			Criteria:      &query.QuestWithTagIdCriteria{Field: "tag_id"},
			Documentation: "Get quests by their tag id",
		},
		{
			Name:          "tag_name",
			Criteria:      &query.QuestWithTagNameCriteria{Field: "tag_name"},
			Documentation: "Get quests by their tag name",
		},
	}
}
