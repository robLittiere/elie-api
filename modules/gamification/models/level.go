package models

type Level struct {
	Id                     int    `json:"id" gorm:"primary_key"`
	Name                   string `json:"name"`
	NextLevelXpRequirement int    `json:"nextLevelXpRequirement"`
	LevelNumber            int    `json:"level_number"`
	CurrencyWon            int    `json:"currencyWon"`
}
