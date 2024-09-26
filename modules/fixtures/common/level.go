package common

import (
	"elie-api/config"
	"elie-api/modules/gamification/models"
)

func CreateBatchLevels() {
	var levels []models.Level

	beginnerLevel := models.Level{
		Name:                   "beginner",
		NextLevelXpRequirement: 10,
		LevelNumber:            1,
		CurrencyWon:            10,
	}

	intermediateLevel := models.Level{
		Name:                   "intermediate",
		NextLevelXpRequirement: 20,
		LevelNumber:            2,
		CurrencyWon:            20,
	}

	advancedLevel := models.Level{
		Name:                   "advanced",
		NextLevelXpRequirement: 30,
		LevelNumber:            3,
		CurrencyWon:            30,
	}

	levels = append(levels, beginnerLevel, intermediateLevel, advancedLevel)
	config.DB.Create(&levels)
}
