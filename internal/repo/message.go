package repo

import (
	"gorm.io/gorm"

	"github.com/mrbelka12000/pdd_tests_bot/internal/models"
)

type MessageRepo struct{}

func NewMessageRepo() *MessageRepo {
	return &MessageRepo{}
}

// CreateMessage inserts a new message
func (r *MessageRepo) CreateMessage(db *gorm.DB, m *models.Message) error {
	return db.Create(&m).Error
}

// GetMessageByID fetches a message by primary key ID
func (r *MessageRepo) GetMessageByID(db *gorm.DB, id int64) (*models.Message, error) {
	var m models.Message
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

// GetMessageByTelegramMessageID fetches a message by TelegramMessageID
func (r *MessageRepo) GetMessageByTelegramMessageID(db *gorm.DB, telegramMessageID int) (*models.Message, error) {
	var m models.Message
	if err := db.Where("telegram_message_id = ?", telegramMessageID).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

// UpdateMessage updates a message record
func (r *MessageRepo) UpdateMessage(db *gorm.DB, m *models.Message) error {
	return db.Save(m).Error
}

// DeleteMessage deletes a message by ID
func (r *MessageRepo) DeleteMessage(db *gorm.DB, id int64) error {
	return db.Delete(&models.Message{}, id).Error
}
