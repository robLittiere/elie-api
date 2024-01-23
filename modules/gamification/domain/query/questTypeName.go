package query

import (
	"elie-api/modules/gamification/models"
	"gorm.io/gorm"
)

type QuestTypeNameCriteria struct {
	Field string
}

func (c *QuestTypeNameCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	var questType models.QuestType
	result := db.Model(&models.QuestType{}).Where("type = ?", value).First(&questType)
	if result.Error != nil {
		return db
	}
	return db.Where("quest_type_id = ?", questType.Id)

}
