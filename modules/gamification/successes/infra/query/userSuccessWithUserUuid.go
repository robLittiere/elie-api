package query

import "gorm.io/gorm"

type UserSuccessWithUserUuidCriteria struct {
	Field string
}

func (u UserSuccessWithUserUuidCriteria) ApplyQuery(db *gorm.DB, value string) *gorm.DB {
	return db.Joins("left join users on users.id = user_successes.user_id").Where("users.uuid = ?", value)
}
