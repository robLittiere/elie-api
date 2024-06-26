package query

import (
	"fmt"
	"gorm.io/gorm"
)

type UserQuestWithQuestIdCriteria struct {
	Field string
}

func (c UserQuestWithQuestIdCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%s = ?", c.Field), value).Session(&gorm.Session{})
}
