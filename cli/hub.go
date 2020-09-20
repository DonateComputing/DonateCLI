package cli

import (
	"flag"
	"fmt"

	"github.com/mfigurski80/DonateCLI/app"
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

	if args := c.fs.Args(); c.fs.NArg() > 0 && args[0] == "ps" {
		recmd := NewHubCommand()
		recmd.Init(args[1:])
		return recmd.Run()
	}

	auth, err := readAuth()
	if err != nil {
		return err
	}

	jobs, err := app.ListHub(*auth, c.isAll, c.isUser)
	if err != nil {
		return err
	}
	for _, j := range jobs {
		fmt.Printf("[%s] %s/%s : %s\n", j.ID[:5], j.Author, j.Title, j.Description)
	}

	return nil
}
