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
	cmd.fs.BoolVar(&cmd.isUser, "u", false, "flag to show existing user jobs")

	return cmd
}

// HubCommand shows all publicly listed jobs and allows filtering by user and purpose
type HubCommand struct {
	fs *flag.FlagSet

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

	args := c.fs.Args()
	if len(args) > 0 && args[0] == "ps" {
		recmd := NewHubCommand()
		recmd.Init(args[1:])
		return recmd.Run()
	}

	auth, err := readAuth()
	if err != nil {
		return err
	}

	jobs, err := app.ListHub(*auth, c.isUser)
	if err != nil {
		return err
	}
	for _, j := range jobs {
		free := "free"
		if j.Runner != "" {
			free = "taken"
		}
		fmt.Printf("\t%s/%s : %s [%s] %s\n", j.Author, j.Title, j.Description, j.OriginalImage, free)
	}
	if len(jobs) <= 0 {
		fmt.Println("Not jobs are currently posted")
	}

	return nil
}
