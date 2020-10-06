package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// StopContainer stops container of given id
func StopContainer(id string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	err = cli.ContainerStop(context.Background(), id, nil)
	if err != nil {
		return err
	}
	return cli.ContainerRemove(context.Background(), id, types.ContainerRemoveOptions{})
}
