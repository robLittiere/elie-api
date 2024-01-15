package models

type User struct {
	Id             int    `json:"id" gorm:"primary_key"`
	Uid            string `json:"uid"`
	Lid            int    `json:"level_id"`
	Level          Level
	Email          string `json:"email" gorm:"unique;not_null"`
	Password       string `json:"password"`
	Username       string `json:"username" gorm:"unique;not_null"`
	Xp             int    `json:"xp"`
	CurrencyAmount int    `json:"currency_amount"`
}
