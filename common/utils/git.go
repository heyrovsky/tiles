package utils

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func CreateOrphanRepository(repo *git.Repository, branchName, appname string) error {
	tree := &object.Tree{}
	treeObj := repo.Storer.NewEncodedObject()
	if err := tree.Encode(treeObj); err != nil {
		return fmt.Errorf("failed to encode empty tree: %w", err)
	}
	treeHash, err := repo.Storer.SetEncodedObject(treeObj)
	if err != nil {
		return fmt.Errorf("failed to store empty tree: %w", err)
	}
	signature := &object.Signature{
		Name:  fmt.Sprintf("%s+%s", appname, GetUsername()),
		Email: fmt.Sprintf("%s+%s@%s", appname, GetUsername(), GetHostName()),
		When:  time.Now(),
	}

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

	refName := plumbing.NewBranchReferenceName(branchName)
	ref := plumbing.NewHashReference(refName, commitHash)
	if err := repo.Storer.SetReference(ref); err != nil {
		return fmt.Errorf("failed to set branch reference: %w", err)
	}

	return nil
}

func IsSSHGitUrl(url string) bool {
	matched, _ := regexp.MatchString(`^git@[\w.-]+:[\w./-]+\.git$`, url)
	return matched
}
