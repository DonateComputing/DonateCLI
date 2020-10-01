package app

import (
	"github.com/mfigurski80/DonateCLI/api"
	"github.com/mfigurski80/DonateCLI/docker"
)

// Start checks out given job id and runs in on local machine
func Start(user string, title string, auth api.AuthStruct) (string, error) {
	ref := api.JobRefStruct{Title: title, User: user}
	job, err := api.GetJob(ref)
	if err != nil {
		return "", err
	}

	// pull n run
	err = docker.PullImage("docker.io/" + job.OriginalImage)
	if err != nil {
		return "", err
	}
	id, err := docker.CreateNewContainer(job.OriginalImage, job.Title)
	if err != nil {
		return id, err
	}

	// mark as being run
	err = api.TakeJob(ref, auth)
	if err != nil {
		// aaah! Dump container we just ran
		err2 := docker.StopContainer(id)
		if err2 != nil {
			return "", err2
		}
		return id, err
	}

	return id, nil
}
