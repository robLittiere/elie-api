package filter

import "elie-api/modules/common/criteria"

// Filters are a list of the queryParam name, the criteria associated to it and a documentation for what is the param
type Filters []Filter

// Filter is a struct that contains the queryParam name, the criteria associated to it and a documentation for what is the param
type Filter struct {
	Name          string
	Criteria      criteria.Criteria
	Documentation string
}
