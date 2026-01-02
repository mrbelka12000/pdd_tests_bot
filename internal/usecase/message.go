package usecase

import (
	"github.com/mrbelka12000/pdd_tests_bot/internal/models"
)

func (uc *UseCase) CreateMessage(obj models.Message) error {
	return uc.messageRepo.CreateMessage(uc.db.DB, &obj)
}

func (uc *UseCase) GetMessageByTelegramMessageID(id int) (*models.Message, error) {
	return uc.messageRepo.GetMessageByTelegramMessageID(uc.db.DB, id)
}

func (uc *UseCase) GetMessageByID(id int64) (*models.Message, error) {
	return uc.messageRepo.GetMessageByID(uc.db.DB, id)
}

func (uc *UseCase) UpdateMessage(m *models.Message) error {
	return uc.messageRepo.UpdateMessage(uc.db.DB, m)
}

func (uc *UseCase) DeleteMessage(id int64) error {
	return uc.messageRepo.DeleteMessage(uc.db.DB, id)
}
