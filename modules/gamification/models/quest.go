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
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type QuestType struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Type string `json:"type"`
}

type UserQuest struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	UserId      int       `json:"user_id" gorm:"primaryKey"`
	QuestId     int       `json:"quest_id" gorm:"primaryKey"`
	Quest       Quest     `json:"quest" gorm:"foreignKey:QuestId"`
	Progression int       `json:"progression" gorm:"default:0"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
