package user

import (
	"elie-api/config"
	"elie-api/modules/user/models"
	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
	"reflect"
)

var f = faker.New()

func CreateUser(data map[string]interface{}) models.User {
	defaults := map[string]interface{}{
		"Email":          f.Internet().Email(),
		"Password":       "rob",
		"Username":       f.Internet().User(),
		"LevelId":        1,
		"Xp":             0,
		"CurrencyAmount": 0,
	}
	user := models.User{}

	val := reflect.ValueOf(&user).Elem()
	for key, _ := range defaults {
		value := data[key]
		if value == nil {
			value = defaults[key]
		}
		val.FieldByName(key).Set(reflect.ValueOf(value))
	}

	user.Uuid = uuid.NewString()
	config.DB.Create(&user)
	return user
}

func GetUserByUuid(uuid string) models.User {
	var user models.User
	config.DB.Where("uuid = ?", uuid).First(&user)
	return user
}

func GetUserByID(id int) models.User {
	var user models.User
	config.DB.First(&user, id)
	return user
}
