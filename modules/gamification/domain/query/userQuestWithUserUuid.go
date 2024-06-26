package query

import "gorm.io/gorm"

type UserQuestWithUserUuid struct {
	Field string
}

func (u UserQuestWithUserUuid) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Joins("left join users on users.id = user_quests.user_id").Where("users.uuid = ?", value).Session(&gorm.Session{})
}
