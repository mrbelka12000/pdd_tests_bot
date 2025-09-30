package usecase

import (
	"github.com/mrbelka12000/pdd_tests_bot/pkg/gorm/postgres"
)

type (
	UseCase struct {
		db       *postgres.Gorm
		userRepo userRepository
		caseRepo caseRepository
		storage  storage
		aiClient aiClient
	}
)

func New(db *postgres.Gorm, ur userRepository, cr caseRepository, st storage, aiClient aiClient) *UseCase {
	return &UseCase{
		db:       db,
		userRepo: ur,
		caseRepo: cr,
		storage:  st,
		aiClient: aiClient,
	}
}
