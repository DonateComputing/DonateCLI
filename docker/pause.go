package docker

import (
	"context"

	"github.com/docker/docker/client"
)

// PauseContainer pauses container with given id
func PauseContainer(id string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	return cli.ContainerPause(context.Background(), id)
}
