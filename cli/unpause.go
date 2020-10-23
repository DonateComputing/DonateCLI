package cli

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mfigurski80/DonateCLI/app"
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
	auth, err := readAuth()
	if err != nil {
		return errors.New("Could not find auth file. Please run `login` command")
	}

	switch c.fs.NArg() {
	case 0:
		err = app.UnpauseAll(*auth)
		if err != nil {
			return err
		}
		fmt.Println("All jobs unpaused")
	case 2:
		err = app.Unpause(c.fs.Arg(0), c.fs.Arg(1), *auth)
		if err != nil {
			return err
		}
		fmt.Printf("Job '%s/%s' unpaused\n", c.fs.Arg(0), c.fs.Arg(1))
	}
	return nil
}
