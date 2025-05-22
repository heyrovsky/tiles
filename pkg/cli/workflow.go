package cli

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/heyrovsky/tiles/common/gitutils"
)

type WorkflowInitCmd struct {
	Name string `arg:"" required:"" help:"Name the tiles workflow."`
}

type WorkflowRemoteCloneCmd struct {
	Url  string `arg:"" required:"" help:"Git URL to clone the workflow from."`
	Name string `help:"Custom name for the workflow folder"`
}

type WorkflowRemoteAddCmd struct {
	Name string `arg:"" required:"" help:"Name of the remote."`
	Url  string `arg:"" required:"" help:"URL of the remote."`
}

type WorkflowRemotePushCmd struct{}

type WorkflowRemoteSyncCmd struct{}

type WorkflowRemoteCmd struct {
	Add   WorkflowRemoteAddCmd   `cmd:"" help:"Add a remote to the workflow."`
	Push  WorkflowRemotePushCmd  `cmd:"" help:"Push workflow to remote."`
	Sync  WorkflowRemoteSyncCmd  `cmd:"" help:"Sync workflow with remote."`
	Clone WorkflowRemoteCloneCmd `cmd:"" help:"Clone a workflow from a git repository."`
}

// functions

func (cmd *WorkflowInitCmd) Run() error {
	fmt.Printf("Initializing new workflow with name: %s\n", cmd.Name)
	_, err := gitutils.InitRepository(cmd.Name)
	if err != nil {
		return err
	}
	return nil
}

func (cmd *WorkflowRemoteSyncCmd) Run() error {
	fmt.Println("Syncing workflow with remote...")
	return nil
}

func (cmd *WorkflowRemotePushCmd) Run() error {
	fmt.Println("Pushing workflow to remote...")
	return nil
}

func (cmd *WorkflowRemoteAddCmd) Run() error {
	fmt.Printf("Adding remote '%s' with URL: %s\n", cmd.Name, cmd.Url)
	return nil
}

func (cmd *WorkflowRemoteCloneCmd) Run(ctx *kong.Context) error {
	fmt.Printf("Cloning workflow from URL: %s\n", cmd.Url)
	_, err := gitutils.CloneRepository(cmd.Url, cmd.Name)
	return err
}
