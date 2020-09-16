package cli

import (
	"flag"
	"fmt"
)

// NewStopCommand creates default stop command
func NewStopCommand() *StopCommand {
	return &StopCommand{
		fs: flag.NewFlagSet("stop", flag.ContinueOnError),
	}
}

// StopCommand stops a currently running job, uploads partial image, and returns it to hub
type StopCommand struct {
	fs *flag.FlagSet
}

// Name returns the name of the stop command (stop)
func (c *StopCommand) Name() string {
	return c.fs.Name()
}

// Init initializes flag data for the stop command
func (c *StopCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

// Run executes the stop command
func (c *StopCommand) Run() error {
	fmt.Printf("Running STOP command. Args are %v\n", c.fs.Args())
	return nil
}
