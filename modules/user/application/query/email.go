package query

import (
	"fmt"
	"gorm.io/gorm"
)

type EmailCriteria struct {
	Field string
}

func (criteria *EmailCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%v = ?", criteria.Field), value)
}
