package models

import "time"

type Quest struct {
	Id             int       `json:"id" gorm:"primary_key"`
	QuestTypeId    int       `json:"quest_type_id"`
	QuestType      QuestType `json:"quest_type" gorm:"foreignKey:QuestTypeId"`
	Name           string    `json:"name"`
	Xp             int       `json:"xp"`
	Difficulty     string    `json:"difficulty"`
	CurrencyReward int       `json:"currency_reward" gorm:"default:0"`
	DoneCondition  int       `json:"done_condition"`
	Tags           string    `json:"tags"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type QuestType struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Type string `json:"type"`
}
