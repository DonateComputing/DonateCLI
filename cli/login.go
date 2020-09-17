package cli

import (
	"errors"
	"flag"
	"fmt"
)

// NewLoginCommand creates a default login command
func NewLoginCommand() *LoginCommand {
	cmd := &LoginCommand{
		fs: flag.NewFlagSet("login", flag.ContinueOnError),
	}
	cmd.fs.BoolVar(&cmd.isRegister, "register", false, "flag to register user if not exists")

	return cmd
}

// LoginCommand creates a login data file and ensures that the account exists on public hub
type LoginCommand struct {
	fs *flag.FlagSet

	isRegister bool
}

// Name returns the name of the login command (login)
func (c *LoginCommand) Name() string {
	return c.fs.Name()
}

// Init initializes flag data for the login command
func (c *LoginCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

// Run executes login command
func (c *LoginCommand) Run() error {
	if c.fs.NArg() != 2 {
		fmt.Println("Usage: login [optional: flags] <username> <password>")
		return errors.New("Bad usage")
	}

	u := c.fs.Arg(0)
	p := c.fs.Arg(1)

	fmt.Printf("Running LOGIN %s %s. Register is '%v'\n", u, p, c.isRegister)
	return nil
}
