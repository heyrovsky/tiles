package workflow

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/heyrovsky/tiles/common/utils"
)

func CloneRepository(url, name, sshKeyLocation, shhKeyPassword string) error {
	targetDir := name
	if targetDir == "" {
		urlParts := strings.Split(url, "/")
		repoName := strings.TrimSuffix(urlParts[len(urlParts)-1], ".git")
		targetDir = repoName
	}
	if err := utils.CreateDirectoryIfNotExists(targetDir); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	cloneOptions := &git.CloneOptions{
		URL:      url,
		Progress: utils.NewDefault(),
	}

	if utils.IsSSHGitUrl(url) {
		auth, err := utils.GetSSHAuthWithPassphrase(sshKeyLocation, shhKeyPassword)
		if err != nil {
			return err
		}

		cloneOptions.Auth = auth
	}

	_, err := git.PlainClone(targetDir, false, cloneOptions)
	if err != nil {
		return err
	}
	return nil
}
