package docker

import (
	"context"

	"github.com/docker/docker/client"
)

// UnpauseContainer unpauses container with given id
func UnpauseContainer(id string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	return cli.ContainerUnpause(context.Background(), id)
}
