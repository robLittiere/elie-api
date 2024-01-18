package query

import (
	"fmt"
	"gorm.io/gorm"
)

type QuizTopicNameCriteria struct {
	Field string
}

func (criteria *QuizTopicNameCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("LOWER(%s->>'name') = LOWER(?)", criteria.Field), value)
}
