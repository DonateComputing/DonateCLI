package cli

import (
	"errors"
	"fmt"
)

func helpCommand() error {
	fmt.Println("Please refer to the repository readme for usage")
	return nil
}

func errorCommand(m string) error {
	fmt.Println(m)
	return errors.New(m)
}

func dneCommand(c string) error {
	return errorCommand(fmt.Sprintf("The command '%s' does not exist", c))
}
