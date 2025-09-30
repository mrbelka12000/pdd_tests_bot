package usecase

import (
	"github.com/mrbelka12000/pdd_tests_bot/internal/models"
)

func (uc *UseCase) CreateUser(obj models.User) error {
	return uc.userRepo.Save(uc.db.DB, obj)
}

func (uc *UseCase) GetAllUsers() ([]models.User, error) {
	return uc.userRepo.GetAllUsers(uc.db.DB)
}

func (uc *UseCase) GetUserByChatID(chatID int64) (models.User, error) {
	return uc.userRepo.GetUserByChatID(uc.db.DB, chatID)
}

func (uc *UseCase) UpdateUser(obj models.User) error {
	return uc.userRepo.UpdateUser(uc.db.DB, obj)
}
