package filters

import (
	"elie-api/modules/common/filter"
	"elie-api/modules/game/domain/query"
)

func GetGameFilters() filter.Filters {
	return filter.Filters{
		{
			Name:          "name",
			Criteria:      &query.GameNameCriteria{Field: "name"},
			Documentation: "Filter by game name",
		},
	}
}
