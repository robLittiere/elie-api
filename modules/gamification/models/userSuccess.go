package models

import (
	"time"
)

type UserSuccess struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	UserId      int       `json:"user_id" gorm:"primaryKey"`
	SuccessId   int       `json:"success_id" gorm:"primaryKey"`
	Success     Success   `json:"success" gorm:"foreignKey:SuccessId"`
	Progression int       `json:"progression" gorm:"default:0"`
	IsCompleted bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserSuccessProgressRequest struct {
	UserUuid  string `json:"user_uuid"`
	SuccessId int    `json:"success_id"`
}
