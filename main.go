package main

import (
	"os"

	"github.com/mfigurski80/DonateCLI/cli"
)

func main() {
	err := cli.Run(os.Args[1:])
	if err != nil {
		panic(err)
	}
}
