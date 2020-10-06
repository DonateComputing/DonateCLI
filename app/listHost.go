package app

import (
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
	for _, c := range dContainers {
		for _, j := range u.Running {
			if c.Names[0] == "/"+j.Title {
				newCont := Container{c, j.User, j.Title}
				containers = append(containers, newCont)
			}
		}
	}

	return containers, nil
}
