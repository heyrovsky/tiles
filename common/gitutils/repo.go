package gitutils

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/heyrovsky/tiles/common/utils"
)

func InitRepository(path string) (*git.Repository, error) {
	if err := utils.CreateDirectoryIfNotExists(path); err != nil {
		return nil, err
	}

	repo, err := git.PlainInit(path, false)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize repository: %w", err)
	}

	return repo, nil
}

func CloneRepository(url string, name string) (*git.Repository, error) {
	// todo : check if its a tiles repo
	targetDir := name

	if targetDir == "" {
		urlParts := strings.Split(url, "/")
		repoName := urlParts[len(urlParts)-1]
		repoName = strings.TrimSuffix(repoName, ".git")
		targetDir = repoName
	}

	if err := utils.CreateDirectoryIfNotExists(targetDir); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}
	cloneOptions := &git.CloneOptions{
		URL:      url,
		Progress: utils.NewDefault(),
	}
	if utils.IsSSHGitUrl(url) {
		auth, err := GetSSHAuth()
		if err != nil {
			return nil, err
		}
		cloneOptions.Auth = auth
	}
	repo, err := git.PlainClone(targetDir, false, cloneOptions)

	if err != nil {
		return nil, fmt.Errorf("failed to clone repository: %w", err)
	}

	return repo, nil

}
