package cli

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mfigurski80/DonateCLI/app"
)

// NewPsCommand creates a default ps command
func NewPsCommand() *PsCommand {
	cmd := &PsCommand{
		fs: flag.NewFlagSet("ps", flag.ContinueOnError),
	}
	cmd.fs.BoolVar(&cmd.isAll, "a", false, "flag to include all completed jobs")

	return cmd
}

// PsCommand displays all images currently being run on computer
type PsCommand struct {
	fs *flag.FlagSet

	isAll bool
}

// Name returns the name of the ps command (ps)
func (c *PsCommand) Name() string {
	return c.fs.Name()
}

// Init initializes flag data for the ps command
func (c *PsCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

// Run executes ps command
func (c *PsCommand) Run() error {
	auth, err := readAuth()
	if err != nil {
		return errors.New("Could not find auth file. Please run `login` command")
	}
	list, err := app.List(*auth, c.isAll)
	if err != nil {
		return err
	}

	if len(list) <= 0 {
		fmt.Println("No jobs are currently running")
	} else {
		for _, item := range list {
			fmt.Printf("[%s] %s/%s %s (%s)\n", item.ID[:5], item.User, item.Title, item.State, item.Status)
		}
	}

	return nil
}
