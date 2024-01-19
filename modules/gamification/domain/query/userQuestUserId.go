package query

import (
	"fmt"
	"gorm.io/gorm"
)

type UserQuestUserIdCriteria struct {
	Field string
}

func (u UserQuestUserIdCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%s = ?", u.Field), value)
}
