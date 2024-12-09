package infrastructure

import (
	"os"
)

type FileLoader struct{}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (l *FileLoader) LoadFiles(paths []string) ([]string, error) {
	var validFiles []string
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			validFiles = append(validFiles, path)
		}
	}
	return validFiles, nil
}
