package models

import "time"

type Quest struct {
	Id             int       `json:"id" gorm:"primary_key"`
	TagId          int       `json:"tag_id"`
	Tag	           Tag       `json:"tag"  gorm:"foreignKey:TagId"`
	Name           string    `json:"name"`
	Xp             int       `json:"xp"`
	Difficulty     string    `json:"difficulty"`
	DoneCondition  int       `json:"done_condition"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
