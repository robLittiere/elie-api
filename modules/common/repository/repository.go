package repository

import (
	"elie-api/modules/common/criteria"
	"elie-api/modules/common/filter"
	"fmt"
	"gorm.io/gorm"
)

type BaseRepo struct {
	DB             *gorm.DB
	FilterRegistry filter.Filters
}

// CriteriaApplicator Interface that base repos will implement
type CriteriaApplicator interface {
	ApplyQuery(db *gorm.DB, value string) *gorm.DB
	BuildQuery(queryParams map[string][]string) error
}

/*
ApplyCriteria will apply the criteria to the query

This is very useful when you want to apply multiple filters to the query`
*/
func (r *BaseRepo) ApplyCriteria(c criteria.Criteria, queryValue string) *gorm.DB {
	return c.ApplyQuery(r.DB, queryValue)
}

/*
BuildQuery will build the query based on the available filters for the resource.

Output: error if the queryparam is not found in the available filters
*/
func (r *BaseRepo) BuildQuery(queryParams map[string][]string) error {

	if len(queryParams) != 0 {
		// parse query parameters
		for key, param := range queryParams {
			for _, value := range param {
				// If there are no available filters for this ressource, return an error
				if len(r.FilterRegistry) == 0 {
					return fmt.Errorf("there are no criterias available for this resource")
				}
				// get the criteria for this specific query parameter
				criteria, err := GetCriteria(key, r.FilterRegistry)
				fmt.Println(key, value)
				if err != nil {
					return err
				}
				r.DB = r.ApplyCriteria(criteria, value)
			}
		}
	}
	return nil
}

func GetCriteria(s string, filterRegistry filter.Filters) (criteria.Criteria, error) {

	// get the filter according to its name
	for _, filter := range filterRegistry {
		if filter.Name == s {
			return filter.Criteria, nil
		}
	}
	return nil, fmt.Errorf("query parameter %v is not a valid filter", s)
}
