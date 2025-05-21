package cli

type WorkflowCmd struct {
	Init   WorkflowInitCmd   `cmd:"" help:"Initialize a new workflow."`
	Clone  WorkflowCloneCmd  `cmd:"" help:"Clone a workflow from a git repository."`
	Remote WorkflowRemoteCmd `cmd:"" help:"Manage workflow remotes."`
}
