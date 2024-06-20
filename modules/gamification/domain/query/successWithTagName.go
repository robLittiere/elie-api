package query

import (
	"gorm.io/gorm"
)

type SuccessWithTagNameCriteria struct {
	Field string
}

func (c *SuccessWithTagNameCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {

	return db.Joins("left join tags on tags.id = successes.tag_id").Where("tags.name = ?", value)
}
