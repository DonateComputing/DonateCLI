package app

import (
	"fmt"

	"github.com/mfigurski80/DonateCLI/api"
	"github.com/mfigurski80/DonateCLI/docker"
)

// Pause pauses given job
func Pause(user string, title string, auth api.AuthStruct) error {
	list, err := List(auth, false)
	if err != nil {
		return err
	}

	for _, c := range list {
		if c.User != user || c.Title != title {
			continue
		}
		return docker.PauseContainer(c.ID)
	}

	return fmt.Errorf("can not find job '%s/%s' on host", user, title)
}

// PauseAll pauses all jobs present
func PauseAll(auth api.AuthStruct) error {
	list, err := List(auth, false)
	if err != nil {
		return err
	}

	for _, c := range list {
		err := docker.PauseContainer(c.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
