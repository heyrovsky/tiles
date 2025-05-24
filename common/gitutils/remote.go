package gitutils

import (
	"fmt"

	"github.com/go-git/go-git/v5/config"
)

// AddRemoteUrltoRepository adds a new remote with the given name and URL.
func AddRemoteUrltoRepository(repoPath, name, url string) error {
	repo, err := openRepo(repoPath)
	if err != nil {
		return err
	}

	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: name,
		URLs: []string{url},
	})
	if err != nil {
		return fmt.Errorf("failed to create remote '%s': %w", name, err)
	}

	fmt.Printf("Remote '%s' added: %s\n", name, url)
	return nil
}

// ShowAllRemotes lists all remotes in the repository.
func ShowAllRemotes(repoPath string) error {
	repo, err := openRepo(repoPath)
	if err != nil {
		return err
	}

	remotes, err := repo.Remotes()
	if err != nil {
		return fmt.Errorf("failed to get remotes: %w", err)
	}

	if len(remotes) == 0 {
		fmt.Println("No remotes found.")
		return nil
	}

	fmt.Println("Git Remotes:")
	for _, remote := range remotes {
		fmt.Printf("- %s:\n", remote.Config().Name)
		for _, url := range remote.Config().URLs {
			fmt.Printf("    %s\n", url)
		}
	}

	return nil
}

// EditRemoteUrl updates the URL of an existing remote.
func EditRemoteUrl(repoPath, name, newUrl string) error {
	repo, err := openRepo(repoPath)
	if err != nil {
		return err
	}

	cfg, err := repo.Config()
	if err != nil {
		return fmt.Errorf("failed to get repo config: %w", err)
	}

	remoteCfg, exists := cfg.Remotes[name]
	if !exists {
		return fmt.Errorf("remote '%s' does not exist", name)
	}

	remoteCfg.URLs = []string{newUrl}
	cfg.Remotes[name] = remoteCfg

	if err := repo.Storer.SetConfig(cfg); err != nil {
		return fmt.Errorf("failed to update remote URL: %w", err)
	}

	fmt.Printf("Remote '%s' updated to: %s\n", name, newUrl)
	return nil
}

// DeleteRemote removes a remote from the repository.
func DeleteRemote(repoPath, name string) error {
	repo, err := openRepo(repoPath)
	if err != nil {
		return err
	}

	if err := repo.DeleteRemote(name); err != nil {
		return fmt.Errorf("failed to delete remote '%s': %w", name, err)
	}

	fmt.Printf("Remote '%s' deleted successfully.\n", name)
	return nil
}
