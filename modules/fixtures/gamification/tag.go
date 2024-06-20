package gamification

import (
	"elie-api/config"
	"elie-api/modules/gamification/models"
	"reflect"
)

func CreateTag(data map[string]interface{}) models.Tag {
	var tag = models.Tag{}

	val := reflect.ValueOf(&tag).Elem()

	defaults := map[string]interface{}{
		"Name": models.PlayGameTag,
	}

	for key, _ := range defaults {
		value := data[key]
		if value == nil {
			value = defaults[key]
		}
		val.FieldByName(key).Set(reflect.ValueOf(value))
	}

	config.DB.Create(&tag)
	return tag
}
