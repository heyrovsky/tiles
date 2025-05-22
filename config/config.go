package config

import (
	"log"
	"strings"

	"github.com/heyrovsky/tiles/common/utils"
)

func LoadConfig(path string) {
	abspath, err := utils.AbsPath(path)
	if err != nil {
		log.Println(err)
	}
	utils.LoadConfig(abspath)
}

func LoadSSHkeyLoaction(path string) {
	abspath, err := utils.AbsPath(path)
	if err != nil {
		log.Println(err)
	}

	SSH_KEY_LOCATION = abspath
}
func LoadSSHkeyPass(password string) {
	SSH_KEY_PASS = strings.TrimSpace(password)
}
