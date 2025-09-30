package usecase

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mrbelka12000/pdd_tests_bot/internal/client/ai"
	"github.com/mrbelka12000/pdd_tests_bot/internal/models"
	"github.com/mrbelka12000/pdd_tests_bot/pkg/pdf"
)

const (
	defaultDir = "data"
)

func (uc *UseCase) Import() error {
	dir, err := os.ReadDir(defaultDir)
	if err != nil {
		return fmt.Errorf("read dir: %w", err)
	}

	pool := make(chan struct{}, 20)

	for _, file := range dir {
		pool <- struct{}{}

		go func(fileName string) {
			defer func() {
				<-pool
			}()

			fullFileName := strings.Join([]string{defaultDir, fileName}, "/")

			text, err := pdf.GetText(fullFileName)
			if err != nil {
				fmt.Println(fmt.Errorf("get text: %w", err))
				return
			}

			info, err := uc.aiClient.GetInfo(ai.InfoRequest{
				Text: text,
			})
			if err != nil {
				fmt.Println(fmt.Errorf("get info: %w", err))
				return
			}

			ptrFileName := fileName
			cs := models.Case{
				Filename:      &ptrFileName,
				Question:      info.Question,
				CorrectAnswer: info.CorrectAnswer,
				CreatedAt:     time.Now().UTC(),
			}

			if err := uc.caseRepo.CreateCase(uc.db.DB, &cs); err != nil {
				fmt.Println(fmt.Errorf("create case: %w", err))
				return
			}

			for _, answer := range info.Answers {
				cs.Answers = append(cs.Answers, models.Answer{
					CaseID: cs.ID,
					Answer: answer.Answer,
					Number: answer.Number,
				})
			}

			if err := uc.caseRepo.UpdateCaseWithAnswers(uc.db.DB, cs); err != nil {
				fmt.Println(fmt.Errorf("update case: %w", err))
				return
			}

			f, err := os.OpenFile(fullFileName, os.O_RDWR, 0600)
			if err != nil {
				fmt.Println(fmt.Errorf("open file: %w", err))
				return
			}

			defer f.Close()

			stat, err := f.Stat()
			if err != nil {
				fmt.Println(fmt.Errorf("stat file: %w", err))
				return
			}

			_, err = uc.storage.UploadFile(context.Background(), f, fileName, "application/pdf", stat.Size())
			if err != nil {
				fmt.Println(fmt.Errorf("upload file: %w", err))
				return
			}

			fmt.Println(fileName, "uploaded")

		}(file.Name())
	}

	return nil
}
