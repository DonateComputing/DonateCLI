package cli

import (
	"flag"
	"fmt"

	"github.com/mfigurski80/DonateCLI/app"
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

	if c.fs.NArg() < 2 {
		return fmt.Errorf("usage: prune <username> <password> ::(auth to dockerhub.io)")
	}

	auth, err := readAuth()
	if err != nil {
		return err
	}

	return app.Prune(*auth, c.fs.Arg(0), c.fs.Arg(1))
}
