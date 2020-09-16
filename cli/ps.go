package cli

import (
	"flag"
	"fmt"
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
	fmt.Printf("Running PS command! Flag -a is '%v'. Args are '%v'\n", c.isAll, c.fs.Args())
	return nil
}
