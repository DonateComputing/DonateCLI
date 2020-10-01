package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// CreateNewContainer creates and runs a container based on given local image
func CreateNewContainer(image string, name string) (string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return "", err
	}

	var gigabyte int64 = 1_073_741_824

	cont, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: image,
		},
		&container.HostConfig{
			Privileged: false,
			Isolation:  "hyperv",
			Resources: container.Resources{
				CPUShares: 4,
				Memory:    gigabyte * 5,
			},
		},
		nil,
		name,
	)
	if err != nil {
		return "", err
	}

	cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	return cont.ID, nil
}
