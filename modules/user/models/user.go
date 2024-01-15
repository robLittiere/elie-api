package models

import (
	"elie-api/modules/gamification/models"
	"time"
)

type User struct {
	Id             int            `json:"id" gorm:"primary_key"`
	Uid            string         `json:"uid"`
	Lid            int            `json:"lid"`
	Level          models.Level   `json:"level" gorm:"foreignKey:Lid"`
	Email          string         `json:"email" gorm:"unique;not_null"`
	Password       string         `json:"password"`
	Username       string         `json:"username" gorm:"unique;not_null"`
	Xp             int            `json:"xp"`
	CurrencyAmount int            `json:"currency_amount"`
	Quests         []models.Quest `json:"quests" gorm:"many2many:user_quests;"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}
