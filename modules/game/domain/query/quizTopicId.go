package query

import (
	"fmt"
	"gorm.io/gorm"
)

type QuizTopicIdCriteria struct {
	Field string
}

func (criteria *QuizTopicIdCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%s->>'id' = ?", criteria.Field), value)
}
