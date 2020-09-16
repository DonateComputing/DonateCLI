package cli

import (
	"flag"
	"fmt"
)

// NewPauseCommand creates default stop command
func NewPauseCommand() *PauseCommand {
	return &PauseCommand{
		fs: flag.NewFlagSet("pause", flag.ContinueOnError),
	}
}

// PauseCommand stops a currently running job, uploads partial image, and returns it to hub
type PauseCommand struct {
	fs *flag.FlagSet
}

// Name returns the name of the pause command (pause)
func (c *PauseCommand) Name() string {
	return c.fs.Name()
}

// Init initializes flag data for the pause command
func (c *PauseCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

// Run executes the stop command
func (c *PauseCommand) Run() error {
	fmt.Printf("Running PAUSE command. Args are %v\n", c.fs.Args())
	return nil
}
