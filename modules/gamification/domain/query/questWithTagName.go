package query

import "gorm.io/gorm"

type QuestWithTagNameCriteria struct {
	Field string
}

func (c *QuestWithTagNameCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Joins("left join tags on tags.id = quests.tag_id").Where("tags.name = ?", value)
}
