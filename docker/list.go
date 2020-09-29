package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ListContainers lists all running containers
func ListContainers() ([]types.Container, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	return containers, nil

	// if len(containers) == 0 {
	// 	fmt.Println("There are no containers running")
	// 	return make([]types.Container, 0), nil
	// }

	// for _, container := range containers {
	// 	fmt.Printf("[%s] %s\n", container.ID, container.Image)
	// }
}
