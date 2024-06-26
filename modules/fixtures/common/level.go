package common

import (
	"elie-api/config"
	"elie-api/modules/gamification/models"
)

func CreateBatchLevels() {
	var levels []models.Level

	basicLevel := models.Level{
		Name:                   "Basic",
		NextLevelXpRequirement: 10,
		LevelNumber:            1,
		CurrencyWon:            10,
	}

	mediumLevel := models.Level{
		Name:                   "Medium",
		NextLevelXpRequirement: 20,
		LevelNumber:            2,
		CurrencyWon:            20,
	}

	advancedLevel := models.Level{
		Name:                   "Advanced",
		NextLevelXpRequirement: 30,
		LevelNumber:            3,
		CurrencyWon:            30,
	}

	levels = append(levels, basicLevel, mediumLevel, advancedLevel)
	config.DB.Create(&levels)
}
