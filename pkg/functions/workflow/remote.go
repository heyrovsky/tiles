package workflow

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

func ShowAllRemote(path string) error {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	remotes, err := repo.Remotes()
	if err != nil {
		return err
	}

	if len(remotes) == 0 {
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

func AddRemoteUrltoLocalRepository(repoPath, name, url string) error {

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}

	_, err = repo.CreateRemote(&config.RemoteConfig{
		Name: name,
		URLs: []string{url},
	})

	if err != nil {
		return nil
	}

	return nil
}

func EditRemoteUrl(repoPath, name, newUrl string) error {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}

	cfg, err := repo.Config()
	if err != nil {
		return err
	}

	remoteCfg, exists := cfg.Remotes[name]
	if !exists {
		return fmt.Errorf("remote '%s' does not exist", name)
	}

	remoteCfg.URLs = []string{newUrl}
	cfg.Remotes[name] = remoteCfg
	if err := repo.Storer.SetConfig(cfg); err != nil {
		return nil
	}

	return nil
}

func DeleteRemote(repoPath, name string) error {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}
	if err := repo.DeleteRemote(name); err != nil {
		return fmt.Errorf("failed to delete remote '%s': %w", name, err)
	}
	return nil
}
