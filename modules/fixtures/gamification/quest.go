package gamification

import (
	"elie-api/config"
	"elie-api/modules/gamification/models"
	userModels "elie-api/modules/user/models"
	"reflect"
)

func CreateQuest(data map[string]interface{}) models.Quest {
	tag := CreateTag()
	var quest = models.Quest{}

	val := reflect.ValueOf(&quest).Elem()

	defaults := map[string]interface{}{
		"Name":           "Win a game",
		"DoneCondition":  3,
		"TagId":         1,
		"Xp":             10,
		"Difficulty":     "easy",
		"CurrencyReward": 0,
	}

	for key, _ := range defaults {
		value := data[key]
		if value == nil {
			value = defaults[key]
		}
		val.FieldByName(key).Set(reflect.ValueOf(value))
	}
	quest.TagId = tag.Id

	config.DB.Create(&quest)
	return quest
}

func CreateTag() models.Tag {
	var tag = models.Tag{
		Id: 1,
		Name: "WinGameTag",
	}
	config.DB.Create(&tag)
	return tag
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
