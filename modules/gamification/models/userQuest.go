package models

import (
	"github.com/google/uuid"
	"time"
)

type UserQuest struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	UserId      int       `json:"user_id" gorm:"primaryKey"`
	QuestId     int       `json:"quest_id" gorm:"primaryKey"`
	Quest       Quest     `json:"quest" gorm:"foreignKey:QuestId"`
	Progression int       `json:"progression" gorm:"default:0"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserQuestProgressRequest struct {
	UserUuid uuid.UUID `json:"user_uuid"`
	QuestId  int       `json:"quest_id"`
}
