package cli

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mfigurski80/DonateCLI/app"
)

// NewPauseCommand creates default pause command
func NewPauseCommand() *PauseCommand {
	return &PauseCommand{
		fs: flag.NewFlagSet("pause", flag.ContinueOnError),
	}
}

// PauseCommand pauses a currently running job without stopping it
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

// Run executes the pause command
func (c *PauseCommand) Run() error {
	auth, err := readAuth()
	if err != nil {
		return errors.New("Could not find auth file. Please run `login` command")
	}

	switch c.fs.NArg() {
	case 0:
		err = app.PauseAll(*auth)
		if err != nil {
			return err
		}
		fmt.Println("All jobs paused")
	case 2:
		err = app.Pause(c.fs.Arg(0), c.fs.Arg(1), *auth)
		if err != nil {
			return err
		}
		fmt.Printf("Job '%s/%s' paused\n", c.fs.Arg(0), c.fs.Arg(1))
	}
	return nil
}
