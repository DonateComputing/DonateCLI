package cli

// Run applies arguments and flags to known list of commands
func Run(args []string) error {
	if len(args) < 1 {
		return helpCommand()
	}

	cmds := []Command{
		NewPsCommand(),
		NewHubCommand(),
		NewStartCommand(),
		NewStopCommand(),
		NewPauseCommand(),
		NewUnpauseCommand(),
		NewPruneCommand(),
		NewAddCommand(),
		NewRmCommand(),
		NewLoginCommand(),
	}

	for _, cmd := range cmds {
		if cmd.Name() == args[0] {
			cmd.Init(args[1:])
			return cmd.Run()
		}
	}

	return helpCommand()
}
