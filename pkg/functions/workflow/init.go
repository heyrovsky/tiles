package workflow

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/heyrovsky/tiles/common/utils"
	"github.com/heyrovsky/tiles/config"
)

func InitRepository(path string) error {
	if err := utils.CreateDirectoryIfNotExists(path); err != nil {
		return fmt.Errorf("failed to create directory for repository: %w", err)
	}

	repo, err := git.PlainInit(path, false)
	if err != nil {
		return fmt.Errorf("failed to initialize git repository: %w", err)
	}

	if err := utils.CreateOrphanRepository(repo, fmt.Sprintf("%s-orchestrator-metadata", config.APP_NAME), config.APP_NAME); err != nil {
		return err
	}
	return nil
}
