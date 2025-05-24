package gitutils

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/heyrovsky/tiles/common/utils"
	"github.com/heyrovsky/tiles/config"
)

func GetRepoPath() (string, error) {
	if config.LOCAL_REPO_LOCATION != "" {
		return config.LOCAL_REPO_LOCATION, nil
	}
	return utils.AbsPath(".")
}

// openRepo opens a git repository at the given path.
func openRepo(repoPath string) (*git.Repository, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("could not open git repo at %s: %w", repoPath, err)
	}
	return repo, nil
}

// createOrphanRepository creates a new orphan branch with an initial empty commit.
// This is useful for initializing a branch without any parent commit.
func createOrphanRepository(repo *git.Repository, branchName string) error {
	// Create an empty tree object
	tree := &object.Tree{}
	treeObj := repo.Storer.NewEncodedObject()
	if err := tree.Encode(treeObj); err != nil {
		return fmt.Errorf("failed to encode empty tree: %w", err)
	}

	treeHash, err := repo.Storer.SetEncodedObject(treeObj)
	if err != nil {
		return fmt.Errorf("failed to store empty tree: %w", err)
	}

	// Prepare author and committer information
	signature := &object.Signature{
		Name:  fmt.Sprintf("%s+%s", config.APP_NAME, utils.GetUsername()),
		Email: fmt.Sprintf("%s+%s@%s", config.APP_NAME, utils.GetUsername(), utils.GetHostname()),
		When:  time.Now(),
	}

	// Create the initial commit with the empty tree
	commit := &object.Commit{
		Author:    *signature,
		Committer: *signature,
		Message:   fmt.Sprintf("Initialize empty branch: %s", branchName),
		TreeHash:  treeHash,
	}

	commitObj := repo.Storer.NewEncodedObject()
	if err := commit.Encode(commitObj); err != nil {
		return fmt.Errorf("failed to encode commit: %w", err)
	}

	commitHash, err := repo.Storer.SetEncodedObject(commitObj)
	if err != nil {
		return fmt.Errorf("failed to store commit: %w", err)
	}

	// Create a new branch reference pointing to the initial commit
	refName := plumbing.NewBranchReferenceName(branchName)
	ref := plumbing.NewHashReference(refName, commitHash)
	if err := repo.Storer.SetReference(ref); err != nil {
		return fmt.Errorf("failed to set branch reference: %w", err)
	}

	return nil
}
