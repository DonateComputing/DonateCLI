package app

import (
	"fmt"

	"github.com/mfigurski80/DonateCLI/api"
	"github.com/mfigurski80/DonateCLI/docker"
)

// Unpause unpauses given job
func Unpause(user string, title string, auth api.AuthStruct) error {
	list, err := List(auth, true)
	if err != nil {
		return err
	}

	for _, c := range list {
		if c.User != user || c.Title != title {
			continue
		}
		return docker.UnpauseContainer(c.ID)
	}

	return fmt.Errorf("can not find job '%s/%s' on host", user, title)
}

// UnpauseAll unpauses all paused jobs
func UnpauseAll(auth api.AuthStruct) error {
	list, err := List(auth, true)
	if err != nil {
		return err
	}

	for _, c := range list {
		if c.State != "paused" {
			continue
		}
		err := docker.UnpauseContainer(c.ID)
		if err != nil {
			return err
		}
	}

	return err
}
