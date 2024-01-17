package query

import (
	"fmt"
	"gorm.io/gorm"
)

type GameNameCriteria struct {
	Field string
}

func (c *GameNameCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("LOWER(%v) = LOWER(?)", c.Field), value)
}
