package main

import (
	"elie-api/cmd/api/modules/user/models"
	"elie-api/cmd/config"
)

func RunSeeders() {
	// Run each seeder here
	seedLevelsIfNeeded()
}

func seedLevelsIfNeeded() {
	var count int64
	config.DB.Model(&models.Level{}).Count(&count)

	if count == 0 {
		// The Level table is empty, so let's seed some data
		levels := []models.Level{
			{
				Name:                   "Beginner",
				NextLevelXpRequirement: 200,
				CurrencyWon:            100,
			},
			{
				Name:                   "Intermediate",
				NextLevelXpRequirement: 400,
				CurrencyWon:            200,
			},
			{
				Name:                   "Advanced",
				NextLevelXpRequirement: 800,
				CurrencyWon:            400,
			},
			{
				Name:                   "Expert",
				NextLevelXpRequirement: 1600,
				CurrencyWon:            800,
			},
		}

		// Batch insert the data into the Level table
		config.DB.Create(&levels)
	}
}
