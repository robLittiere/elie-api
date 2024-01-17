package query

import (
	"fmt"
	"gorm.io/gorm"
)

type QuizTopicCriteria struct {
	Field string
}

func (criteria *QuizTopicCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%v = ?", criteria.Field), value)
}
