package gamification

import (
	"elie-api/config"
	"elie-api/modules/gamification/models"
	"reflect"
)

func CreateSuccess(data map[string]interface{}) models.Success {
	tag := CreateTag(map[string]interface{}{})
	var success = models.Success{}

	val := reflect.ValueOf(&success).Elem()

	defaults := map[string]interface{}{
		"Name":            "Win a game",
		"Xp":              10,
		"TagId":           tag.Id,
		"DoneCondition":   3,
		"ProgressionRank": 0,
		"CurrencyReward":  0,
	}

	for key, _ := range defaults {
		value := data[key]
		if value == nil {
			value = defaults[key]
		}
		val.FieldByName(key).Set(reflect.ValueOf(value))
	}

	config.DB.Create(&success)
	return success

}
