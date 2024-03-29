package query

import (
	"fmt"
	"gorm.io/gorm"
)

type UserSuccessUserIdCriteria struct {
	Field string
}

func (u UserSuccessUserIdCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%s = ?", u.Field), value)
}
