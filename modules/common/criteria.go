package common

import "gorm.io/gorm"

type Criteria interface {
	ApplyQuery(db *gorm.DB, value string) *gorm.DB
}
