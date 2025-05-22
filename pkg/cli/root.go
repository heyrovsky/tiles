package cli

import (
	"log"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/heyrovsky/tiles/config"
)

// CLI defines the root command structure.
var CLI struct {
	Workflow WorkflowCmd `cmd:"" help:"Manage workflow operations."`
	Version  VersionCmd  `cmd:"" help:"Show application version."`

	SSHKey     string `help:"Path to SSH private key file (for auth)" defaults:""`
	SSHKeyPass string `help:"Password to the ssh key" defaults:""`
}

// Run parses and executes CLI commands.
func Run() {
	ctx := kong.Parse(&CLI,
		kong.Name(config.APP_NAME),
		kong.Description("Tiles – A Git-based orchestrator for infrastructure and automation."),
		kong.Vars{"version": config.APP_VERSION},
		kong.UsageOnError(),
	)

	if strings.TrimSpace(CLI.SSHKey) != "" {
		config.LoadSSHkeyLoaction(strings.TrimSpace(CLI.SSHKey))
	}

	if CLI.SSHKeyPass != "" {
		config.LoadSSHkeyPass(CLI.SSHKeyPass)
	}

	if err := ctx.Run(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}
