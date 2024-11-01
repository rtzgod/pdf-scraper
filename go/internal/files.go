package internal

import (
	"fmt"
	"os"
)

func DirFiles(path string) ([]os.FileInfo, error) {
	const op = "internal.files.DirFiles"

	dir, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	return files, nil
}
