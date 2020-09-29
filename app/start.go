package app

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/mfigurski80/DonateCLI/api"
)

// Start checks out given job id and runs in on local machine
func Start(id string, auth api.AuthStruct) error {

	job, err := api.GetJob(id)
	if err != nil {
		return err
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	closer, err := cli.ImagePull(context.Background(), "docker.io/"+job.OriginalImage, types.ImagePullOptions{})
	defer closer.Close()

	return err
}
