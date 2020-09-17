package main

import (
	"fmt"

	"github.com/mfigurski80/DonateCLI/api"
)

func main() {
	// err := cli.Run(os.Args[1:])
	// if err != nil {
	// 	panic(err)
	// }

	auth := *api.MakeAuthStruct("miko", "pass")

	fmt.Println(api.GetUser(auth))
}
