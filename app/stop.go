package app

import (
	"fmt"

	"github.com/mfigurski80/DonateCLI/api"
	"github.com/mfigurski80/DonateCLI/docker"
)

// Stop returns given job and stops its execution
func Stop(user string, title string, auth api.AuthStruct) error {
	ref := api.JobRefStruct{Title: title, User: user}

	list, err := List(auth, true)
	if err != nil {
		return err
	}
	for _, c := range list { // find in currently run...
		if c.User != user || c.Title != title {
			continue
		}
		err = docker.StopContainer(c.ID)
		if err != nil {
			return err
		}
		err := api.ReturnJob(ref, api.JobReturnStruct{}, auth)
		if err != nil {
			return err
		}
	}

	return fmt.Errorf("can not find job '%s/%s' on host", user, title)
}

// StopAll returns all given jobs and stops all execution
func StopAll(auth api.AuthStruct) error {
	list, err := List(auth, true)
	if err != nil {
		return err
	}
	for _, c := range list {
		err := docker.StopContainer(c.ID)
		if err != nil {
			return err
		}
		ref := api.JobRefStruct{User: c.User, Title: c.Title}
		err = api.ReturnJob(ref, api.JobReturnStruct{}, auth)
		if err != nil {
			return err
		}
	}

	return nil
}
