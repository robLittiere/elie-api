package query

import (
	"fmt"
	"gorm.io/gorm"
)

type QuizTopicIdArrayCriteria struct {
	Field string
}

func (criteria *QuizTopicIdArrayCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	// Convert string '1,2,3' to gorm usable []string{'1','2','3'}
	list := []string{}
	for _, v := range value {
		list = append(list, string(v))
	}

	return db.Where(fmt.Sprintf("%s->>'id' IN (?)", criteria.Field), list)
}
