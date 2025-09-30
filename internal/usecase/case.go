package usecase

import "github.com/mrbelka12000/pdd_tests_bot/internal/models"

func (uc *UseCase) GetCase(id int64) (*models.Case, error) {
	return uc.caseRepo.GetCaseByID(uc.db.DB, id)
}
