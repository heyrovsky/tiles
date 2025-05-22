package cli

type WorkflowCmd struct {
	Init   WorkflowInitCmd   `cmd:"" help:"Initialize a new workflow."`
	Remote WorkflowRemoteCmd `cmd:"" help:"Manage workflow remotes."`
}
