package infrastructure

import (
	"elie-api/modules/common/repository"
	"elie-api/modules/user/models"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	repository.BaseRepo
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{BaseRepo: repository.BaseRepo{DB: db}}
}

func (r *UserRepo) Find() ([]models.User, error) {
	var users []models.User
	result := r.DB.Preload("Level").Preload("Quests").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepo) BuildQueryAndFind(queryParams map[string][]string) ([]models.User, error) {
	var users []models.User
	err := r.BuildQuery(queryParams)
	if err != nil {
		return nil, err
	}
	users, err = r.Find()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) CreateUser(user *models.User) error {
	if err := r.IsUserInDb(*user); err != nil {
		return err
	}

	user.Uuid = uuid.New().String()
	// Add basic level
	lid := 1
	result := r.DB.Table("levels").Select("id").Where("name = ?", "Basic").Scan(&lid)
	if result.Error != nil {
		return result.Error
	}
	user.Lid = lid

	if err := user.HashPassword(user.Password, 2); err != nil {
		return err
	}

	result = r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepo) IsUserInDb(user models.User) error {
	result := r.DB.Select("id").Where("email = ?", user.Email).Limit(1).Find(&models.User{})
	if result.RowsAffected != 0 {
		err := fmt.Errorf("user already registered with this email")
		return err
	}
	return nil
}
