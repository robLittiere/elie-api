package query

import (
	"fmt"
	"gorm.io/gorm"
)

type QuestWithIdCriteria struct {
	Field string
}

func (c QuestWithIdCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%s = ?", c.Field), value)
}
