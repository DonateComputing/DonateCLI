package cli

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mfigurski80/DonateCLI/app"
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
	if c.fs.NArg() < 2 {
		return errors.New("User and Title of public runnable job is required")
	}
	auth, err := readAuth()
	if err != nil {
		return errors.New("Could not find auth file. Please run `login` command")
	}

	username := c.fs.Arg(0)
	title := c.fs.Arg(0)
	id, err := app.Start(username, title, *auth)

	fmt.Printf("Started '%s/%s' as [%s]", username, title, id)

	return nil
}
