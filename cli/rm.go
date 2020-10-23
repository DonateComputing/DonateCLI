package cli

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mfigurski80/DonateCLI/app"
)

// NewRmCommand creates a default rm command
func NewRmCommand() *RmCommand {
	return &RmCommand{
		fs: flag.NewFlagSet("rm", flag.ContinueOnError),
	}
}

// RmCommand removes image of given id from public hub
type RmCommand struct {
	fs *flag.FlagSet
}

// Name returns the name of the rm command (rm)
func (c *RmCommand) Name() string {
	return c.fs.Name()
}

// Init initializes flag data for the rm command
func (c *RmCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

// Run executes rm command
func (c *RmCommand) Run() error {
	if c.fs.NArg() < 1 {
		return errors.New("Id of owned public job is required")
	}

	auth, err := readAuth()
	if err != nil {
		return errors.New("Could not find auth file. Please run `login` command")
	}

	fmt.Printf("Running RM %v command!\n", c.fs.Arg(0))
	return app.Remove(*auth, c.fs.Arg(0))
}
