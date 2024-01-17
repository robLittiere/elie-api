package models

import "time"

type Game struct {
	Id               int       `json:"id" gorm:"primary_key"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	CatchPhrase      string    `json:"catch_phrase"`
	CanBeMultiplayer bool      `json:"can_be_multiplayer"`
	GameVersion      string    `json:"game_version"`
	CreatedAt        time.Time `json:"createdAt"`
	ReleasedAt       time.Time `json:"releasedAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
