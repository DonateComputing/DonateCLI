package cli

import (
	"errors"
	"flag"
	"fmt"
)

// NewAddCommand creates a default add command
func NewAddCommand() *AddCommand {
	cmd := &AddCommand{
		fs: flag.NewFlagSet("add", flag.ContinueOnError),
	}
	cmd.fs.BoolVar(&cmd.allowMultiple, "allow-multiple", false, "flag to allow multiple runners of job")

	return cmd
}

// AddCommand creates and uploads job to public hub
type AddCommand struct {
	fs *flag.FlagSet

	allowMultiple bool
}

// Name returns the name of the add command (add)
func (c *AddCommand) Name() string {
	return c.fs.Name()
}

// Init initializes flag data for the add command
func (c *AddCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

// Run executes the add command
func (c *AddCommand) Run() error {

	switch c.fs.NArg() {
	case 3:
		fmt.Printf("Running the ADD %v command! Flag --allow-multiple is %v\n", c.fs.Args(), c.allowMultiple)
	case 0:
		fmt.Printf("Running the ADD command! Flag --allow-multiple is '%v'. Args are %v\n", c.allowMultiple, c.fs.Args())
	default:
		return errors.New("Incorrect number of arguments")
	}
	return nil
}
