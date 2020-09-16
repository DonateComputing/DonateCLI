package cli

// Runner interface for commands
type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}
