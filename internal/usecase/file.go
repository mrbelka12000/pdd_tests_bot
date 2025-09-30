package usecase

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mrbelka12000/pdd_tests_bot/pkg/image"
)

const (
	extractedDir = "extracted"
)

func (uc *UseCase) DownloadFile(fileName string) (string, error) {
	// Get the file reader from storage
	reader, err := uc.storage.DownloadFile(context.Background(), fileName)
	if err != nil {
		return "", err
	}

	// Create local file
	out, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	// Copy the content from storage reader to local file
	_, err = io.Copy(out, reader)
	if err != nil {
		return "", err
	}

	if err := image.Extract(fileName); err != nil {
		return "", err
	}

	if err := os.Remove(fileName); err != nil {
		return "", err
	}

	dir, err := os.ReadDir(extractedDir)
	if err != nil {
		return "", err
	}

	for _, f := range dir {
		if strings.Contains(f.Name(), fileName[:5]) {
			return filepath.Join(extractedDir, f.Name()), nil
		}
	}

	return "", errors.New("file not found")
}
