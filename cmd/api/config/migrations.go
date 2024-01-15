package config

import (
	models2 "elie-api/cmd/api/modules/user/models"
)

func Migrate() {
	err := DB.AutoMigrate(
		&models2.User{},
		&models2.Level{},
	)
	if err != nil {
		return
	}

	// Add other migrations here if necessary
}

// DropTables Drop tables if they exist
func DropTables() {
	// Example
	/*	if (DB.Migrator().HasTable(&models.User{})) {
			DB.Migrator().DropTable(&models.User{})
		}

	*/
}
