package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func LoadConfig(path string) {
	viper.SetConfigFile(path)
	viper.SetConfigType("toml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}
}

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
