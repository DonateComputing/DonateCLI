package cli

import (
	"fmt"
)

func helpCommand() error {
	fmt.Println("Please refer to the repository readme for usage")
	return nil
}

func errorCommand(m string) error {
	fmt.Println(m)
	return fmt.Errorf(m)
}

func dneCommand(c string) error {
	return errorCommand(fmt.Sprintf("The command '%s' does not exist", c))
}
