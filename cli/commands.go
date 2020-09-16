package cli

import (
	"fmt"
)

func helpCommand() error {
	fmt.Println("Please refer to the repository readme for usage")
	return nil
}
