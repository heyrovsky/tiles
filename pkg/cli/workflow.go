package cli

import (
	"fmt"
)

type WorkflowInitCmd struct {
	Name string `arg:"" required:"" help:"Name the tiles workflow."`
}

func (cmd *WorkflowInitCmd) Run() error {
	fmt.Printf("Initializing new workflow with name: %s\n", cmd.Name)
	return nil
}

type WorkflowCloneCmd struct {
	Url string `arg:"" required:"" help:"Git URL to clone the workflow from."`
}

func (cmd *WorkflowCloneCmd) Run() error {
	fmt.Printf("Cloning workflow from URL: %s\n", cmd.Url)
	return nil
}

type WorkflowRemoteAddCmd struct {
	Name string `arg:"" required:"" help:"Name of the remote."`
	Url  string `arg:"" required:"" help:"URL of the remote."`
}

func (cmd *WorkflowRemoteAddCmd) Run() error {
	fmt.Printf("Adding remote '%s' with URL: %s\n", cmd.Name, cmd.Url)
	return nil
}

type WorkflowRemotePushCmd struct{}

func (cmd *WorkflowRemotePushCmd) Run() error {
	fmt.Println("Pushing workflow to remote...")
	return nil
}

type WorkflowRemoteSyncCmd struct{}

func (cmd *WorkflowRemoteSyncCmd) Run() error {
	fmt.Println("Syncing workflow with remote...")
	return nil
}

type WorkflowRemoteCmd struct {
	Add  WorkflowRemoteAddCmd  `cmd:"" help:"Add a remote to the workflow."`
	Push WorkflowRemotePushCmd `cmd:"" help:"Push workflow to remote."`
	Sync WorkflowRemoteSyncCmd `cmd:"" help:"Sync workflow with remote."`
}
