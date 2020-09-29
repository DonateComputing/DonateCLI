package docker

import (
	"context"

	"github.com/docker/docker/client"
)

// StopContainer stops container of given id
func StopContainer(id string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	err = cli.ContainerStop(context.Background(), id, nil)
	return err
}
