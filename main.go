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

	// auth := *api.MakeAuthStruct("miko", "pass")
	// job := *api.MakePostJobStruct("Pi", "Computing PI!", "NA", false)

	// fmt.Println(api.GetUser(auth))
	// fmt.Println(api.AddJob(auth, job))
	// fmt.Println(api.GetJobs())
	// fmt.Println(api.DeleteJob(auth, "1600318152628776859"))
}
