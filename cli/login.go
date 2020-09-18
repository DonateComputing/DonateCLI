package cli

import (
	"errors"
	"flag"
	"fmt"

	"github.com/mfigurski80/DonateCLI/api"
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

	// parse data
	auth := api.MakeAuthStruct(c.fs.Arg(0), c.fs.Arg(1))

	// verify login or register
	if c.isRegister {
		err := api.Register(*auth)
		if err != nil {
			fmt.Printf("Failed to register user '%v'\n", auth.Username)
			return err
		}
	} else {
		_, err := api.GetUser(*auth)
		if err != nil {
			fmt.Printf("Failed to login as user '%v'\n", auth.Username)
			return err
		}
	}

	// write to settings file
	err := writeSettings(auth.Username, auth.Password)
	if err != nil {
		return err
	}

	fmt.Printf("Login as '%v' successful... settings file updated\n", auth.Username)
	return nil
}
