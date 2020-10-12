package app

import (
	"fmt"
	"strings"

	"github.com/mfigurski80/DonateCLI/api"
	"github.com/mfigurski80/DonateCLI/docker"
)

// Prune removes and returns all completed jobs
func Prune(auth api.AuthStruct, username string, password string) error {
	list, err := List(auth, true)
	if err != nil {
		return err
	}

	for _, c := range list {
		if c.State != "exited" {
			continue
		}
		newTag := "docker.io/" + username + "/" + strings.ToLower(c.Title) + ":completed"
		imgID, err := docker.ContainerCommit(c.ID, newTag)
		if err != nil {
			panic(err)
		}
		err = docker.PushImage(username, password, newTag)
		if err != nil {
			return err
		}

		jobRef := api.JobRefStruct{Title: c.Title, User: c.User}
		jobData := api.JobReturnStruct{Image: newTag}
		err = api.ReturnJob(jobRef, jobData, auth)
		if err != nil {
			return err
		}

		fmt.Println(imgID)
	}
	return nil
}
