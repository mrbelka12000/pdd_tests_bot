package repo

import (
	"gorm.io/gorm"

	"github.com/mrbelka12000/pdd_tests_bot/internal/models"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (u UserRepo) Save(db *gorm.DB, user models.User) error {
	return db.Save(&user).Error
}

func (u UserRepo) GetUserByID(db *gorm.DB, id int64) (*models.User, error) {
	var user models.User
	err := db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepo) GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserRepo) GetUserByChatID(db *gorm.DB, chatID int64) (models.User, error) {
	var user models.User
	err := db.Where("chat_id = ?", chatID).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u UserRepo) UpdateUser(db *gorm.DB, user models.User) error {
	return db.Updates(user).Error
}

func (u UserRepo) DeleteUser(db *gorm.DB, id int64) error {
	return db.Delete(&models.User{}, id).Error
}
