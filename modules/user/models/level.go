package models

type Level struct {
	ID                     int    `json:"id" gorm:"primary_key"`
	Name                   string `json:"name"`
	NextLevelXpRequirement int    `json:"nextLevelXpRequirement"`
	CurrencyWon            int    `json:"currencyWon"`
}
