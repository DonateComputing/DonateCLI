package cli

import (
	"flag"
	"fmt"
)

// NewHubCommand creates a default hub command
func NewHubCommand() *HubCommand {
	cmd := &HubCommand{
		fs: flag.NewFlagSet("hub", flag.ContinueOnError),
	}
	cmd.fs.BoolVar(&cmd.isAll, "a", false, "flag to include all taken and completed jobs")
	cmd.fs.BoolVar(&cmd.isUser, "u", false, "flag to show only user jobs")

	return cmd
}

// HubCommand shows all publicly listed jobs and allows filtering by user and purpose
type HubCommand struct {
	fs *flag.FlagSet

	isAll  bool
	isUser bool
}

// Name returns the name of the hub command (hub)
func (c *HubCommand) Name() string {
	return c.fs.Name()
}

// Init initializes flag data for the hub command
func (c *HubCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

// Run executes hub command
func (c *HubCommand) Run() error {
	fmt.Printf("Running HUB command! Flag -a is '%v', flag -u is '%v'. Args are '%v'\n", c.isAll, c.isUser, c.fs.Args())
	return nil
}
