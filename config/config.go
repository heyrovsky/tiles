package config

import (
	"log"

	"github.com/heyrovsky/tiles/common/utils"
)

func LoadConfig(path string) {
	abspath, err := utils.AbsPath(path)
	if err != nil {
		log.Println(err)
	}
	utils.LoadConfig(abspath)
}
