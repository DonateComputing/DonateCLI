package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// PushImage pushes given image to public docker registry
func PushImage(username string, password string, name string) error {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	jsonBytes, err := json.Marshal(types.AuthConfig{Username: username, Password: password})
	if err != nil {
		return err
	}
	authStr := base64.URLEncoding.EncodeToString(jsonBytes)

	pusher, err := cli.ImagePush(ctx, name, types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		return err
	}

	defer pusher.Close()
	// io.Copy(os.Stdout, pusher)

	return nil
}
