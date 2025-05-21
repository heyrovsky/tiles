package cli

import (
	"fmt"

	"github.com/heyrovsky/tiles/config"
)

type VersionCmd struct {
}

func (version *VersionCmd) Run() error {
	fmt.Println("Tiles version : ", config.APP_VERSION)
	fmt.Println("Commit SHA    : ", config.APP_GIT_HASH)
	fmt.Println("go version    : ", config.APP_GO_VERSION)
	return nil
}
