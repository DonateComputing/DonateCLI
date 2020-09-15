package main

// ActionFunc takes a command and args and performs an action
type ActionFunc func() error

// Command is small container for cli actionable command
type Command struct {
	Name        string
	Description string
	FlagList    []string
	SubCommands map[string]Command
	Action      ActionFunc
}
