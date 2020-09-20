package cli

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/mfigurski80/DonateCLI/api"
)

var homeDir, _ = os.UserHomeDir()
var configPath = homeDir + "/.donateconfig"

func writeAuth(auth api.AuthStruct) error {
	file, err := json.MarshalIndent(auth, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configPath, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func readAuth() (*api.AuthStruct, error) {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var settings api.AuthStruct
	err = json.Unmarshal(file, &settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}
