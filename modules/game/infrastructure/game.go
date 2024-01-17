package infrastructure

import (
	"elie-api/modules/common/repository"
	"elie-api/modules/game/application/filters"
	"elie-api/modules/game/models"
	"gorm.io/gorm"
)

type GameRepo struct {
	repository.BaseRepo
}

func NewGameRepo(db *gorm.DB) *GameRepo {
	return &GameRepo{BaseRepo: repository.BaseRepo{DB: db, FilterRegistry: filters.GetGameFilters()}}
}

func (repo *GameRepo) Find() ([]models.Game, error) {
	var games []models.Game
	result := repo.DB.Find(&games)
	if result.Error != nil {
		return nil, result.Error
	}
	return games, nil
}

func (repo *GameRepo) BuildQueryAndFind(queryParams map[string][]string) ([]models.Game, error) {
	var games []models.Game
	err := repo.BuildQuery(queryParams)
	if err != nil {
		return nil, err
	}
	games, err = repo.Find()
	if err != nil {
		return nil, err
	}

	return games, nil
}
