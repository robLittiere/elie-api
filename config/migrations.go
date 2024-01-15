package config

import "elie-api/modules/user/models"

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Level{},
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
