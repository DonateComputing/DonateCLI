package cli

import (
	"flag"
	"fmt"
)

// NewPruneCommand creates default prune command
func NewPruneCommand() *PruneCommand {
	return &PruneCommand{
		fs: flag.NewFlagSet("prune", flag.ContinueOnError),
	}
}

// PruneCommand removes all jobs that have completed
type PruneCommand struct {
	fs *flag.FlagSet
}

// Name returns the name of the prune command (prune)
func (c *PruneCommand) Name() string {
	return c.fs.Name()
}

// Init initializes flag data for the prune command
func (c *PruneCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

// Run executes the prune command
func (c *PruneCommand) Run() error {
	fmt.Printf("Running PRUNE command. Args are %v\n", c.fs.Args())
	return nil
}
