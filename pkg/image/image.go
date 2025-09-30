package image

import (
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

const (
	extractedDir = "extracted"
)

func Extract(filename string) error {
	if err := os.MkdirAll(extractedDir, 0o755); err != nil {
		return fmt.Errorf("create extracted dir: %w", err)
	}

	var selectedPages []string = nil

	if err := api.ExtractImagesFile(filename, extractedDir, selectedPages, nil); err != nil {
		return fmt.Errorf("extract image: %w", err)
	}

	return nil
}
