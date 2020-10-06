package cli

import (
	"flag"
	"fmt"

	"github.com/mfigurski80/DonateCLI/app"
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
	auth, err := readAuth()
	if err != nil {
		return err
	}

	switch c.fs.NArg() {
	case 2:
		err = app.Stop(c.fs.Arg(0), c.fs.Arg(1), *auth)
		if err != nil {
			return err
		}
		fmt.Printf("Job '%s/%s' stopped and returned\n", c.fs.Arg(0), c.fs.Arg(1))
	case 0:
		err = app.StopAll(*auth)
		if err != nil {
			return err
		}
		fmt.Println("All jobs stopped and returned")
	default:
		fmt.Println("Usage: stop <username> <title>")
	}
	return nil
}
