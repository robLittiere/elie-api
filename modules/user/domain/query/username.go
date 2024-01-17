package query

import (
	"fmt"
	"gorm.io/gorm"
)

type UsernameCriteria struct {
	Field string
}

func (criteria *UsernameCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%v = ?", criteria.Field), value)
}
