package main

// ActionFunc takes a command and args and performs an action
type ActionFunc func([]string, []string) error

// Command is small container for cli actionable command
type Command struct {
	Name        string
	Description string
	FlagList    []string
	SubCommands map[string]Command
	Action      ActionFunc
}

// NewCommand returns a new command with some preformed data
func NewCommand(name string, action ActionFunc) *Command {
	return &Command{
		Name:        name,
		Description: name,
		Action:      action,
	}
}

// Run runs the (presumably root) command with given args
func Run(c Command, inp []string) error {
	args, flags := parseArgsFlags(inp[1:])
	return c.Action(args, flags)
}

func parseArgsFlags(inp []string) ([]string, []string) {
	args := make([]string, 0)
	flags := make([]string, 0)
	for _, s := range inp {
		if s[0] == '-' {
			flags = append(flags, s[1:])
			continue
		}
		args = append(args, s)
	}
	return args, flags
}
