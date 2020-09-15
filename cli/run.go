package cli

// Run applies arguments and flags to known list of commands
func Run(args []string) error {
	if len(args) < 2 {
		return helpCommand()
	}

	switch args[1] {

	case "help":
		return helpCommand()

	case "-h":
		return helpCommand()

	default:
		return dneCommand(args[1])

	}

}
