package models

import (
	"elie-api/modules/gamification/models"
	"time"
)

type Success struct {
	Id              int        `json:"id" gorm:"primary_key"`
	Name            string     `json:"name"`
	Xp              int        `json:"xp"`
	TagId           int        `json:"tag_id"`
	Tag             models.Tag `json:"tag"  gorm:"foreignKey:TagId"`
	ParentId        *int       `json:"parent_id" gorm:"default:null"`
	ParentSuccess   *Success   `json:"parent_success" gorm:"foreignKey:ParentId"`
	DoneCondition   int        `json:"done_condition"`
	ProgressionRank int        `json:"progression_rank"`
	CurrencyReward  int        `json:"currency_reward"`
	ShortName       string     `json:"short_name"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
