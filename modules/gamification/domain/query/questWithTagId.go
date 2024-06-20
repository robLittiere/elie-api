package query

import "gorm.io/gorm"

type QuestWithTagIdCriteria struct {
	Field string
}

func (c *QuestWithTagIdCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Where("tag_id = ?", value)
}
