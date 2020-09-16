package cli

import (
	"flag"
	"fmt"
)

// NewUnpauseCommand creates default unpause command
func NewUnpauseCommand() *UnpauseCommand {
	return &UnpauseCommand{
		fs: flag.NewFlagSet("unpause", flag.ContinueOnError),
	}
}

// UnpauseCommand unpauses previously paused job without restarting it
type UnpauseCommand struct {
	fs *flag.FlagSet
}

// Name returns the name of the unpause command (unpause)
func (c *UnpauseCommand) Name() string {
	return c.fs.Name()
}

// Init initializes flag data for the unpause command
func (c *UnpauseCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

// Run executes the upause command
func (c *UnpauseCommand) Run() error {
	fmt.Printf("Running UNPAUSE command. Args are %v\n", c.fs.Args())
	return nil
}
