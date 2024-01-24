package config

import (
	gameModels "elie-api/modules/game/models"
	gamificationModels "elie-api/modules/gamification/models"
	userModels "elie-api/modules/user/models"
)

func Migrate() {
	err := DB.AutoMigrate(
		&userModels.User{},
		&gameModels.Game{},
		&gameModels.QuizGame{},
		&gameModels.UserQuiz{},
		&gamificationModels.Level{},
		&gamificationModels.Quest{},
		&gamificationModels.UserQuest{},
		&gamificationModels.Success{},
	)
	if err != nil {
		return
	}

	// Setup join tables
	err = DB.SetupJoinTable(&userModels.User{}, "Quests", &gamificationModels.UserQuest{})
	if err != nil {
		return
	}

	// Add other migrations here if you want them played automatically
	// Otherwise, use goose to make migrations in the migrations folder
}

// DropTables Drop tables if they exist
func DropTables() {
	// Example
	/*	if (DB.Migrator().HasTable(&models.User{})) {
			DB.Migrator().DropTable(&models.User{})
		}

	*/
}
