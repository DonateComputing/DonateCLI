package app

import (
	"github.com/docker/docker/api/types"
	"github.com/mfigurski80/DonateCLI/api"
	"github.com/mfigurski80/DonateCLI/docker"
)

// List returns list of all containers
func List(auth api.AuthStruct, all bool) ([]types.Container, error) {
	u, err := api.GetUser(auth)
	if err != nil {
		return []types.Container{}, nil
	}

	containers, err := docker.ListContainers(all)
	if err != nil {
		return containers, nil
	}

	containers = filterContainers(containers, func(c types.Container) bool {
		for _, j := range u.Running {
			if c.Names[0] == "/"+j.Title {
				return true
			}
		}
		return false
	})

	return containers, nil
}
