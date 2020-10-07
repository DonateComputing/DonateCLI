package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ContainerCommit converts container to image
func ContainerCommit(id string, name string) (string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return "", err
	}

	res, err := cli.ContainerCommit(context.Background(), id, types.ContainerCommitOptions{Reference: "dockerhub.io/" + name})
	if err != nil {
		return "", err
	}

	newTag := "dockerhub.io/" + name + ":completed"
	cli.ImageTag(context.Background(), res.ID, newTag)

	return res.ID, nil
}
