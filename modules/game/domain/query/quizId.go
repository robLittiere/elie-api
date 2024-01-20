package query

import (
	"fmt"
	"gorm.io/gorm"
)

type QuizIdCriteria struct {
	Field string
}

func (c *QuizIdCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%s->>'id' = ?", c.Field), value)
}
