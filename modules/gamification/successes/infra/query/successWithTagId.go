package query

import (
	"fmt"
	"gorm.io/gorm"
)

type SuccessWithTagIdCriteria struct {
	Field string
}

func (c *SuccessWithTagIdCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%s = ?", c.Field), value)

}
