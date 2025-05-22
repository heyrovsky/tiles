package utils

import (
	"fmt"
	"regexp"
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

func IsSSHGitUrl(url string) bool {

	matched, _ := regexp.MatchString(`^git@[\w.-]+:[\w./-]+\.git$`, url)
	return matched
}
