package query

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type QuizTopicIdArrayCriteria struct {
	Field string
}

func (criteria *QuizTopicIdArrayCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	list := strings.Split(value, ",")
	return db.Where(fmt.Sprintf("%s->>'id' IN (?)", criteria.Field), list)
}
