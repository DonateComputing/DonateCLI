package cli

import (
	"fmt"
)

func helpCommand() error {
	fmt.Println("Usage: app [cmd] [flags] [args]")
	fmt.Println("Refer to repository readme for further usage details")
	return nil
}
