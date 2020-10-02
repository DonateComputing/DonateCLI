package cli

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mfigurski80/DonateCLI/app"
)

// NewAddCommand creates a default add command
func NewAddCommand() *AddCommand {
	cmd := &AddCommand{
		fs: flag.NewFlagSet("add", flag.ContinueOnError),
	}
	return cmd
}

// AddCommand creates and uploads job to public hub
type AddCommand struct {
	fs *flag.FlagSet
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
	auth, err := readAuth()
	if err != nil {
		return err
	}
	args := c.fs.Args()

	switch c.fs.NArg() {
	case 3:
		err := app.Add(args[0], args[1], args[2], *auth)
		if err != nil {
			return err
		}
	case 0:
		// TODO: interface
	default:
		return errors.New("Usage: add <title> <description> <image>")
	}

	fmt.Printf("Successfully added job '%s'\n", args[0])
	return nil
}
