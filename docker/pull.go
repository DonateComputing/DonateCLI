package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// PullImage pulls image from remote registry
func PullImage(image string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	closer, err := cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
	defer closer.Close()

	return err
}
