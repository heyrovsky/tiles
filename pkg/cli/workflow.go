package cli

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/heyrovsky/tiles/config"
	"github.com/heyrovsky/tiles/pkg/functions/workflow"
)

// ==== Structs ====

type WorkflowInitCmd struct {
	Name string `arg:"" required:"" help:"Name the tiles workflow."`
}

type WorkflowRemoteAddCmd struct {
	Name string `arg:"" required:"" help:"Name of the remote."`
	Url  string `arg:"" required:"" help:"URL of the remote."`
}

type WorkflowRemoteEditCmd struct {
	Name string `arg:"" required:"" help:"Name of the remote."`
	Url  string `arg:"" required:"" help:"URL of the remote."`
}

type WorkflowRemoteDeleteCmd struct {
	Name string `arg:"" required:"" help:"Name of the remote."`
}

type WorkflowRemoteShowCmd struct{}

type WorkflowRemotePushCmd struct{}

type WorkflowRemoteSyncCmd struct{}

type WorkflowRemoteCloneCmd struct {
	Url  string `arg:"" required:"" help:"Git URL to clone the workflow from."`
	Name string `help:"Custom name for the workflow folder"`
}

type WorkflowRemoteCmd struct {
	Add    WorkflowRemoteAddCmd    `cmd:"" help:"Add a remote to the workflow."`
	Show   WorkflowRemoteShowCmd   `cmd:"" help:"Show's all remote of the workflow"`
	Push   WorkflowRemotePushCmd   `cmd:"" help:"Push workflow to remote."`
	Sync   WorkflowRemoteSyncCmd   `cmd:"" help:"Sync workflow with remote."`
	Clone  WorkflowRemoteCloneCmd  `cmd:"" help:"Clone a workflow from a git repository."`
	Edit   WorkflowRemoteEditCmd   `cmd:"" help:"Edit a remote URL of the workflow."`
	Delete WorkflowRemoteDeleteCmd `cmd:"" help:"Delete a remote from the workflow."`
}

// ==== Functions ====

func (cmd *WorkflowInitCmd) Run() error {
	fmt.Printf("Initializing new workflow with name: %s\n", cmd.Name)
	return workflow.InitRepository(cmd.Name)
}

func (cmd *WorkflowRemoteShowCmd) Run() error {
	repoPath, err := config.GetLocalRepositoryLocation()
	if err != nil {
		return fmt.Errorf("failed to determine repository path: %w", err)
	}

	if err := workflow.ShowAllRemote(repoPath); err != nil {
		return fmt.Errorf("failed to show remotes: %w", err)
	}

	return nil
}

func (cmd *WorkflowRemoteAddCmd) Run() error {
	fmt.Printf("Adding remote '%s' with URL: %s\n", cmd.Name, cmd.Url)

	repoPath, err := config.GetLocalRepositoryLocation()
	if err != nil {
		return fmt.Errorf("failed to determine repository path: %w", err)
	}

	if err := workflow.AddRemoteUrltoLocalRepository(repoPath, cmd.Name, cmd.Url); err != nil {
		return fmt.Errorf("failed to add remote: %w", err)
	}
	return nil
}

func (cmd *WorkflowRemoteEditCmd) Run() error {
	fmt.Printf("Editing remote '%s' to new URL: %s\n", cmd.Name, cmd.Url)

	repoPath, err := config.GetLocalRepositoryLocation()
	if err != nil {
		return fmt.Errorf("failed to determine repository path: %w", err)
	}

	if err := workflow.EditRemoteUrl(repoPath, cmd.Name, cmd.Url); err != nil {
		return fmt.Errorf("failed to edit remote: %w", err)
	}
	return nil
}

func (cmd *WorkflowRemoteDeleteCmd) Run() error {
	fmt.Printf("Deleting remote '%s'\n", cmd.Name)

	repoPath, err := config.GetLocalRepositoryLocation()
	if err != nil {
		return fmt.Errorf("failed to determine repository path: %w", err)
	}

	if err := workflow.DeleteRemote(repoPath, cmd.Name); err != nil {
		return fmt.Errorf("failed to delete remote: %w", err)
	}
	return nil
}

func (cmd *WorkflowRemotePushCmd) Run() error {
	fmt.Println("Pushing workflow to remote...")
	return nil
}

func (cmd *WorkflowRemoteSyncCmd) Run() error {
	fmt.Println("Syncing workflow with remote...")
	return nil
}

func (cmd *WorkflowRemoteCloneCmd) Run(ctx *kong.Context) error {
	fmt.Printf("Cloning workflow from URL: %s\n", cmd.Url)

	return workflow.CloneRepository(cmd.Url, cmd.Name, config.SSH_KEY_LOCATION, config.SSH_KEY_PASS)
}
