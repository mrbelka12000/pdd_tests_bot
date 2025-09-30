package repo

import (
	"gorm.io/gorm"

	"github.com/mrbelka12000/pdd_tests_bot/internal/models"
)

type CaseRepo struct {
}

func NewCaseRepo() *CaseRepo {
	return &CaseRepo{}
}

func (r *CaseRepo) CreateCase(db *gorm.DB, c *models.Case) error {
	return db.Create(&c).Error
}

func (r *CaseRepo) GetCaseByID(db *gorm.DB, id int64) (*models.Case, error) {
	var c models.Case
	err := db.Preload("Answers").First(&c, id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CaseRepo) GetAllCases(db *gorm.DB) ([]models.Case, error) {
	var cases []models.Case
	err := db.Preload("answers").Find(&cases).Error
	return cases, err
}

func (r *CaseRepo) UpdateCase(db *gorm.DB, c models.Case) error {
	updatesMap := make(map[string]interface{})

	updatesMap["filename"] = c.Filename
	updatesMap["question"] = c.Question
	updatesMap["correct_answer"] = c.CorrectAnswer

	return db.Model(&models.Case{}).
		Where("id = ?", c.ID).
		Updates(updatesMap).Error
}

func (r *CaseRepo) UpdateCaseWithAnswers(db *gorm.DB, c models.Case) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := r.UpdateCase(tx, c); err != nil {
			return err
		}

		if err := tx.Where("case_id = ?", c.ID).Delete(&models.Answer{}).Error; err != nil {
			return err
		}

		for _, ans := range c.Answers {
			ans.CaseID = c.ID
			if err := tx.Create(&ans).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *CaseRepo) DeleteCase(db *gorm.DB, id int64) error {
	return db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Where("case_id = ?", id).Delete(&models.Answer{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.Case{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *CaseRepo) GetRandomCase(db *gorm.DB) (*models.Case, error) {
	var cs models.Case
	err := db.Preload("Answers").Order("RANDOM()").Limit(1).Find(&cs).Error
	if err != nil {
		return nil, err
	}
	return &cs, nil
}
