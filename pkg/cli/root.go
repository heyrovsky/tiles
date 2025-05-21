package cli

import (
	"log"

	"github.com/alecthomas/kong"
	"github.com/heyrovsky/tiles/config"
)

// CLI defines the root command structure.
var CLI struct {
	Workflow WorkflowCmd `cmd:"" help:"Manage workflow operations."`
	Version  VersionCmd  `cmd:"" help:"Show application version."`
}

// Run parses and executes CLI commands.
func Run() {
	ctx := kong.Parse(&CLI,
		kong.Name(config.APP_NAME),
		kong.Description("Tiles – A Git-based orchestrator for infrastructure and automation."),
		kong.Vars{"version": config.APP_VERSION},
		kong.UsageOnError(),
	)

	if err := ctx.Run(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
