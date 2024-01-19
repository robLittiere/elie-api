package models

type Success struct {
	Id            int    `json:"id" gorm:"primary_key"`
	Name          string `json:"name"`
	Xp            int    `json:"xp"`
	Difficulty    string `json:"difficulty"`
	DoneCondition int    `json:"done_condition"`
	CurrencyWon   int    `json:"currency_won"`
}
