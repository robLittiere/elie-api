package main

import (
	"context"
	"database/sql"
	"github.com/elie-prbl/elie-api/api/config"
	"github.com/elie-prbl/elie-api/api/modules/user/models"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddLevelsToLevelTableGo, downAddLevelsToLevelTableGo)
}

func upAddLevelsToLevelTableGo(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	seedLevelsIfNeeded()
	return nil
}

func downAddLevelsToLevelTableGo(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	config.DB.Migrator().DropTable(&models.Level{})
	config.DB.AutoMigrate(&models.Level{})
	return nil
}
