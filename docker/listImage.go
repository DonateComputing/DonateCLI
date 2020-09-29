package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ListImages lists all images installed on local machine
func ListImages() ([]types.ImageSummary, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	return images, err
}
