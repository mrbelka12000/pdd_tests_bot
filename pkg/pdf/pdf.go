package pdf

import (
	"strings"

	"github.com/ledongthuc/pdf"
)

func GetText(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var b strings.Builder
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		page := r.Page(pageIndex)
		if page.V.IsNull() {
			continue
		}

		rows, err := page.GetTextByRow()
		if err != nil {
			return "", err
		}

		for _, row := range rows {
			for _, mark := range row.Content {
				// mark.S — текстовый фрагмент как есть (с пунктуацией)
				b.WriteString(mark.S)
			}
			b.WriteByte('\n')
		}
		b.WriteByte('\n')
	}

	return b.String(), nil
}
