package cli

// Command interface for commands
type Command interface {
	Init([]string) error
	Run() error
	Name() string
}
