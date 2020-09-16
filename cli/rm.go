package cli

import (
	"errors"
	"flag"
	"fmt"
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

	fmt.Printf("Running RM %v command!\n", c.fs.Arg(0))

	return nil
}
