package usecase

import (
	"github.com/mrbelka12000/pdd_tests_bot/internal/models"
)

func (uc *UseCase) GetRandomCase() (*models.Case, error) {
	return uc.caseRepo.GetRandomCase(uc.db.DB)
}
