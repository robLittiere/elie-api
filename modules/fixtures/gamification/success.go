package gamification

import (
	"elie-api/config"
	"elie-api/modules/gamification/models"
	userModels "elie-api/modules/user/models"
	"reflect"
)

func CreateSuccess(data map[string]interface{}) models.Success {
	var success = models.Success{}

	val := reflect.ValueOf(&success).Elem()

	defaults := map[string]interface{}{
		"Name":            "Win a game",
		"Xp":              10,
		"DoneCondition":   3,
		"ProgressionRank": 0,
		"CurrencyReward":  0,
	}

	// Add tagId or create default tag
	tagId := 0
	if data["TagId"] != nil {
		tagId = data["TagId"].(int)
	} else {
		tagId = CreateTag(map[string]interface{}{}).Id
	}
	defaults["TagId"] = tagId

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

func CreateUserSuccess(user userModels.User, success models.Success, data map[string]interface{}) models.UserSuccess {
	var userSuccess = models.UserSuccess{}
	val := reflect.ValueOf(&userSuccess).Elem()

	defaults := map[string]interface{}{
		"Progression": 0,
		"IsCompleted": false,
	}
	for key, _ := range defaults {
		value := data[key]
		if value == nil {
			value = defaults[key]
		}
		val.FieldByName(key).Set(reflect.ValueOf(value))
	}
	userSuccess.UserId = user.Id
	userSuccess.SuccessId = success.Id

	config.DB.Create(&userSuccess)
	return userSuccess
}

func GetUserSuccess(sid int, uid int) models.UserSuccess {
	var userSuccess models.UserSuccess
	config.DB.Preload("Success").Where("success_id = ? AND user_id = ?", sid, uid).First(&userSuccess)
	return userSuccess
}

func GetUserSuccesses(uid int) []models.UserSuccess {
	var userSuccesses []models.UserSuccess
	config.DB.Preload("Success").Where("user_id = ?", uid).Find(&userSuccesses)
	return userSuccesses
}
