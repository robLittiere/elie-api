package models

import "time"

type Game struct {
	Id               int       `json:"id" gorm:"primary_key"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	CatchPhrase      string    `json:"catch_phrase"`
	CanBeMultiplayer bool      `json:"can_be_multiplayer"`
	Theme            string    `json:"theme" gorm:"default:'any'"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type GameData struct {
	Id   int                    `json:"id" gorm:"primary_key"`
	Gid  int                    `json:"game_id"`
	Game Game                   `json:"game" gorm:"foreignKey:Gid"`
	Data map[string]interface{} `json:"data" gorm:"type:jsonb"`
}
