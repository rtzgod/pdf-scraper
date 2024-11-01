package internal

import (
	"bytes"
	"fmt"
	"github.com/ledongthuc/pdf"
)

func ReadPDF(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", fmt.Errorf("ошибка при открытии PDF: %v", err)
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", fmt.Errorf("unable to get plain text")
	}

	if _, err := buf.ReadFrom(b); err != nil {
		return "", fmt.Errorf("ошибка при чтении текста из PDF: %v", err)
	}

	return buf.String(), nil
}
