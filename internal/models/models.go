package models

import "time"

type (
	Case struct {
		ID            int64     `gorm:"primary_key,AUTO_INCREMENT"`
		Filename      *string   `gorm:"column:filename"`
		Question      string    `gorm:"column:question"`
		CorrectAnswer int       `gorm:"column:correct_answer"`
		CreatedAt     time.Time `gorm:"column:created_at"`
		Answers       []Answer  `gorm:"foreignKey:case_id;references:id"` // <-- important
	}

	Answer struct {
		CaseID int64  `gorm:"column:case_id;primaryKey"`
		Number int    `gorm:"column:number;primaryKey"`
		Answer string `gorm:"column:answer"`
	}
)

func (Case) TableName() string {
	return "cases"
}

func (Answer) TableName() string {
	return "answers"
}
