package app

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/mfigurski80/DonateCLI/api"
	"github.com/mfigurski80/DonateCLI/docker"
)

// Container is struct that holds all container data
type Container struct {
	types.Container
	User  string
	Title string
}

// List returns list of all containers
func List(auth api.AuthStruct, all bool) ([]Container, error) {
	u, err := api.GetUser(auth)
	if err != nil {
		return nil, nil
	}

	dContainers, err := docker.ListContainers(all)
	if err != nil {
		return nil, nil
	}

	containers := make([]Container, 0)
OUTER:
	for _, j := range u.Running {
		for _, c := range dContainers {
			if c.Names[0] == "/"+j.Title {
				newCont := Container{c, j.User, j.Title}
				containers = append(containers, newCont)
				continue OUTER
			}
		}
		if !all {
			continue
		}
		// not found running? Update
		err := api.ReturnJob(j, api.JobReturnStruct{}, auth)
		if err != nil {
			return nil, fmt.Errorf("Host list out of sync with Hub -- failed to correct containers: %v", err)
		}
	}

	return containers, nil
}
