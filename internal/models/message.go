package models

type (
	Message struct {
		ID                int64 `gorm:"primary_key"`
		ChatID            int64 `gorm:"column:chat_id"`
		TelegramMessageID int   `gorm:"column:telegram_message_id"`
		CaseID            int64 `gorm:"column:case_id"`
	}
)
