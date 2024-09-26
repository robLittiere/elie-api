package service

import (
	infrastructure2 "elie-api/modules/gamification/infrastructure"
	"elie-api/modules/user/infrastructure"
	"gorm.io/gorm"
)

type UserSuccessProgressService struct {
	UserRepo        *infrastructure.UserRepo
	UserSuccessRepo *infrastructure2.UserSuccessRepo
}

func NewUserSuccessProgressService(conn *gorm.DB) *UserSuccessProgressService {
	return &UserSuccessProgressService{
		UserRepo:        infrastructure.NewUserRepo(conn),
		UserSuccessRepo: infrastructure2.NewUserSuccessRepo(conn),
	}
}
