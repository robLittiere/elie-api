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
		&gamificationModels.UserSuccess{},
		&gamificationModels.Success{},
		&gamificationModels.Tag{},
	)
	if err != nil {
		return
	}

	// Setup join tables
	err = DB.SetupJoinTable(&userModels.User{}, "Quests", &gamificationModels.UserQuest{})
	if err != nil {
		return
	}

	err = DB.SetupJoinTable(&userModels.User{}, "Success", &gamificationModels.UserSuccess{})
	if err != nil {
		return
	}

	// Add other migrations here if you want them played automatically
	// Otherwise, use goose to make migrations in the migrations folder
}

// DropTables Drop tables if they exist
func DropTables() {
	// Example
	if (DB.Migrator().HasTable(userModels.User{})) {
		DB.Migrator().DropTable(userModels.User{})
	}
	if (DB.Migrator().HasTable(gameModels.Game{})) {
		DB.Migrator().DropTable(gameModels.Game{})
	}
	if (DB.Migrator().HasTable(gameModels.QuizGame{})) {
		DB.Migrator().DropTable(gameModels.QuizGame{})
	}
	if (DB.Migrator().HasTable(gameModels.UserQuiz{})) {
		DB.Migrator().DropTable(gameModels.UserQuiz{})
	}
	if (DB.Migrator().HasTable(gamificationModels.Level{})) {
		DB.Migrator().DropTable(gamificationModels.Level{})
	}
	if (DB.Migrator().HasTable(gamificationModels.Quest{})) {
		DB.Migrator().DropTable(gamificationModels.Quest{})
	}
	if (DB.Migrator().HasTable(gamificationModels.UserQuest{})) {
		DB.Migrator().DropTable(gamificationModels.UserQuest{})
	}
	if (DB.Migrator().HasTable(gamificationModels.Tag{})) {
		DB.Migrator().DropTable(gamificationModels.Tag{})
	}
	if (DB.Migrator().HasTable(gamificationModels.UserSuccess{})) {
		DB.Migrator().DropTable(gamificationModels.UserSuccess{})
	}
	if (DB.Migrator().HasTable(gamificationModels.Success{})) {
		DB.Migrator().DropTable(gamificationModels.Success{})
	}
}
