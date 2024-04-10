package models

import (
	"time"
)

type Success struct {
	Id              int       `json:"id" gorm:"primary_key"`
	Name            string    `json:"name"`
	Xp              int       `json:"xp"`
	Difficulty      string    `json:"difficulty"`
	DoneCondition   int       `json:"done_condition"`
	ProgressionRank int       `json:"progression_rank"`
	CurrencyReward  int       `json:"currency_reward"`
	Tags            string    `json:"tags"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
