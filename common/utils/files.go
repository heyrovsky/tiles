package utils

import (
	"fmt"
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
