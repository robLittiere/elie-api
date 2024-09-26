package gamification

import (
	"elie-api/config"
	"elie-api/modules/gamification/models"
	userModels "elie-api/modules/user/models"
	"reflect"
)

func CreateQuest(data map[string]interface{}) models.Quest {
	var quest = models.Quest{}

	val := reflect.ValueOf(&quest).Elem()

	defaults := map[string]interface{}{
		"Name":          "Win a game",
		"DoneCondition": 3,
		"Xp":            10,
		"Difficulty":    "easy",
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

	config.DB.Create(&quest)
	return quest
}

func CreateUserQuest(user userModels.User, quest models.Quest) models.UserQuest {
	var userQuest = models.UserQuest{
		UserId:      user.Id,
		QuestId:     quest.Id,
		Progression: 0,
		IsCompleted: false,
	}
	config.DB.Create(&userQuest)
	return userQuest
}

func GetUserQuest(qid int, uid int) models.UserQuest {
	var userQuest models.UserQuest
	config.DB.Preload("Quest").Where("quest_id = ? AND user_id = ?", qid, uid).First(&userQuest)
	return userQuest
}
