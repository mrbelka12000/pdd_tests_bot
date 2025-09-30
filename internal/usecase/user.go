package usecase

import "github.com/mrbelka12000/pdd_tests_bot/internal/models"

func (uc *UseCase) CreateUser(obj models.User) error {
	return uc.userRepo.CreateUser(uc.db.DB, obj)
}

func (uc *UseCase) GetAllUsers() ([]models.User, error) {
	return uc.userRepo.GetAllUsers(uc.db.DB)
}
