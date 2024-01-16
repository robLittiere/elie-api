package query

import (
	"fmt"
	"gorm.io/gorm"
)

type UuidCriteria struct {
	Field string
}

func (criteria *UuidCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%v = ?", criteria.Field), value)
}
