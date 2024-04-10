package gamification

import (
	"elie-api/config"
	"elie-api/modules/gamification/models"
	"reflect"
)

func CreateBasicLevel() {
	level := models.Level{
		Name:                   "Basic",
		NextLevelXpRequirement: 10,
		LevelNumber:            1,
		CurrencyWon:            10,
	}
	config.DB.Create(&level)
}

func CreateLevel(data map[string]interface{}) models.Level {
	var level = models.Level{}

	defaults := map[string]interface{}{
		"Name":                   "Basic",
		"NextLevelXpRequirement": 10,
		"LevelNumber":            2,
		"CurrencyWon":            20,
	}
	val := reflect.ValueOf(&level).Elem()

	for key, _ := range defaults {
		value := data[key]
		if value == nil {
			value = defaults[key]
		}
		val.FieldByName(key).Set(reflect.ValueOf(value))
	}

	config.DB.Create(&level)
	return level
}
