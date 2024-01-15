package config

func Migrate() {
	err := DB.AutoMigrate()
	if err != nil {
		return
	}
}

// DropTables Drop tables if they exist
func DropTables() {
	// Example
	/*	if (DB.Migrator().HasTable(&models.User{})) {
			DB.Migrator().DropTable(&models.User{})
		}

	*/
}
