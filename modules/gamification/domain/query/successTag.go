package query

import (
	"elie-api/modules/gamification/models"
	"gorm.io/gorm"
)

type SuccessTagCriteria struct {
	Field string
}

func (c *SuccessTagCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	var tag models.Tag
	if c.Field == "tag_id" {
		result := db.Model(&models.Tag{}).Where("id = ?", value).First(&tag)
		if result.Error != nil {
			return db
		}
	} else if c.Field == "tag" {
		result := db.Model(&models.Tag{}).Where("name = ?", value).First(&tag)
		if result.Error != nil {
			return db
		}
	}
	return db.Where("tag_id = ?", tag.Id)
}
