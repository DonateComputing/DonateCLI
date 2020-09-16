package cli

import (
	"errors"
	"flag"
	"fmt"
)

// NewStartCommand creates a default start command
func NewStartCommand() *StartCommand {
	return &StartCommand{
		fs: flag.NewFlagSet("start", flag.ContinueOnError),
	}
}

// StartCommand pulls and begins running a public job with the given id
type StartCommand struct {
	fs *flag.FlagSet
}

// Name returns the name of the start command (start)
func (c *StartCommand) Name() string {
	return c.fs.Name()
}

// Init initializes flag data for the start command
func (c *StartCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

// Run executes the start command
func (c *StartCommand) Run() error {
	if c.fs.NArg() < 1 {
		return errors.New("ID of public runnable job is required")
	}

	fmt.Printf("Running START '%v' command!", c.fs.Arg(0))
	return nil
}
