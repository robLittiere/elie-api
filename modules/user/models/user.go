package models

import (
	"elie-api/modules/gamification/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id             int                `json:"id" gorm:"primary_key"`
	Uuid           string             `json:"uuid"`
	LevelId        int                `json:"level_id"`
	Level          models.Level       `json:"level" gorm:"foreignKey:LevelId"`
	Email          string             `json:"email" gorm:"unique;not_null"`
	Password       string             `json:"password"`
	Username       string             `json:"username" gorm:"unique;not_null"`
	Xp             int                `json:"xp" gore:"default:0"`
	CurrencyAmount int                `json:"currency_amount" gorm:"default:0"`
	UserQuests     []models.UserQuest `json:"quests" gorm:"foreignKey:UserId"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
}

type PublicUser struct {
	Uuid           string `json:"uuid" `
	Lid            int    `json:"lid"`
	Level          models.Level
	Email          string `json:"email"`
	Username       string `json:"username"`
	UserQuests     []models.UserQuest
	Xp             int `json:"xp"`
	CurrencyAmount int `json:"currency_amount"`
}

func (user *User) Serialize() PublicUser {
	return PublicUser{
		Uuid:           user.Uuid,
		Lid:            user.LevelId,
		Level:          user.Level,
		Email:          user.Email,
		Username:       user.Username,
		UserQuests:     user.UserQuests,
		Xp:             user.Xp,
		CurrencyAmount: user.CurrencyAmount,
	}
}

func (user *User) HashPassword(password string, cost int) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
