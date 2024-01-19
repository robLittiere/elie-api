package models

import "time"

type Quest struct {
	Id            int       `json:"id" gorm:"primary_key"`
	Qtid          int       `json:"qtid"`
	QuestType     QuestType `json:"quest_type" gorm:"foreignKey:Qtid"`
	Name          string    `json:"name"`
	Xp            int       `json:"xp"`
	Difficulty    int       `json:"difficulty"`
	CurrencyWon   int       `json:"currency_won" gorm:"default:0"`
	DoneCondition string    `json:"done_condition"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type QuestType struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Type string `json:"type"`
}

type UserQuest struct {
	Id        int `gorm:"primaryKey"`
	UserId    int `gorm:"primaryKey"`
	QuestId   int `gorm:"primaryKey"`
	CreatedAt time.Time
}
