package repository

import (
	"elie-api/modules/common/criteria"
	"elie-api/modules/user/application/filters"
	"fmt"
	"gorm.io/gorm"
)

type BaseRepo struct {
	DB *gorm.DB
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
func (r *BaseRepo) ApplyCriteria(c criteria.Criteria, queryValue string) {
	r.DB = c.ApplyQuery(r.DB, queryValue)
}

/*
BuildQuery will build the query based on the available filters for the resource.

Output: error if the queryparam is not found in the available filters
*/
func (r *BaseRepo) BuildQuery(queryParams map[string][]string) error {

	if len(queryParams) != 0 {
		for key, param := range queryParams {
			for _, value := range param {
				criteria, err := filters.GetUserFilter(key)
				fmt.Println(key, value)
				if err != nil {
					return err
				}
				r.ApplyCriteria(criteria, value)
			}
		}
	}
	return nil
}
