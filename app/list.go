package app

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// List returns list of all containers
func List(all bool) ([]types.Container, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: all})
	if err != nil {
		return nil, err
	}

	containers = filter(containers, func(c types.Container) bool {
		if strings.HasSuffix(c.Image, "donate-api") {
			return true
		}
		return false
	})

	return containers, nil
}
