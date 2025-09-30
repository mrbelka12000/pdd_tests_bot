package usecase

import (
	"context"
	"io"

	"gorm.io/gorm"

	"github.com/mrbelka12000/pdd_tests_bot/internal/client/ai"
	"github.com/mrbelka12000/pdd_tests_bot/internal/models"
)

type (
	userRepository interface {
		CreateUser(db *gorm.DB, user models.User) error
		GetUserByID(db *gorm.DB, id int64) (*models.User, error)
		GetAllUsers(db *gorm.DB) ([]models.User, error)
		GetUserByChatID(db *gorm.DB, chatID string) (*models.User, error)
		UpdateUser(db *gorm.DB, user models.User) error
		DeleteUser(db *gorm.DB, id int64) error
	}

	caseRepository interface {
		CreateCase(db *gorm.DB, c *models.Case) error
		GetCaseByID(db *gorm.DB, id int64) (*models.Case, error)
		GetAllCases(db *gorm.DB) ([]models.Case, error)
		UpdateCase(db *gorm.DB, c models.Case) error
		UpdateCaseWithAnswers(db *gorm.DB, c models.Case) error
		DeleteCase(db *gorm.DB, id int64) error
		GetRandomCase(db *gorm.DB) (*models.Case, error)
	}

	storage interface {
		UploadFile(ctx context.Context, file io.Reader, objectName, contentType string, fileSize int64) (string, error)
		DownloadFile(ctx context.Context, objectName string) (io.Reader, error)
	}

	aiClient interface {
		GetInfo(req ai.InfoRequest) (*ai.InfoResponse, error)
	}
)
