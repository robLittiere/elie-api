package query

import (
	"fmt"
	"gorm.io/gorm"
)

type QuizTitleCriteria struct {
	Field string
}

func (c *QuizTitleCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("LOWER(%s->>'title') LIKE LOWER(?)", c.Field), "%"+value+"%")
}
