package gitutils

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/heyrovsky/tiles/common/utils"
)

// InitRepository initializes a new Git repository at the given path.
// It ensures the directory exists, initializes the repo, and writes orchestrator-specific metadata to the local Git config.
func InitRepository(path string) (*git.Repository, error) {
	if err := utils.CreateDirectoryIfNotExists(path); err != nil {
		return nil, fmt.Errorf("failed to create directory for repository: %w", err)
	}

	repo, err := git.PlainInit(path, false)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize git repository: %w", err)
	}

	return repo, nil
}

// CloneRepository clones a Git repository from the given URL into a local directory.
// If `name` is empty, the repository name is inferred from the URL.
// SSH authentication is handled automatically if the URL is an SSH remote.
func CloneRepository(url, name string) (*git.Repository, error) {
	targetDir := name
	if targetDir == "" {
		urlParts := strings.Split(url, "/")
		repoName := strings.TrimSuffix(urlParts[len(urlParts)-1], ".git")
		targetDir = repoName
	}

	if err := utils.CreateDirectoryIfNotExists(targetDir); err != nil {
		return nil, fmt.Errorf("failed to create target directory: %w", err)
	}

	cloneOptions := &git.CloneOptions{
		URL:      url,
		Progress: utils.NewDefault(),
	}

	if utils.IsSSHGitUrl(url) {
		auth, err := GetSSHAuth()
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve SSH auth: %w", err)
		}
		cloneOptions.Auth = auth
	}

	repo, err := git.PlainClone(targetDir, false, cloneOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to clone repository from %s: %w", url, err)
	}

	return repo, nil
}
