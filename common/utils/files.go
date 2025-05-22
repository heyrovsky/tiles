package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// AbsPath returns the absolute path of the given folder.
func AbsPath(folder string) (string, error) {
	absPath, err := filepath.Abs(folder)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}
	return absPath, nil
}

func CreateDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	return nil
}
