package query

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type QuizTopicNameArrayCriteria struct {
	Field string
}

func (c QuizTopicNameArrayCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	list := strings.Split(value, ",")
	return db.Where(fmt.Sprintf("LOWER(%v->>'name') IN (?)", c.Field), list)
}
