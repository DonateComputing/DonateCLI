package cli

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var homeDir, _ = os.UserHomeDir()
var configPath = homeDir + "/.donateconfig"

func makeSettingsStruct(u string, p string) *settingsStruct {
	return &settingsStruct{
		Username: u,
		Password: p,
	}
}

type settingsStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func writeSettings(username string, password string) error {
	data := makeSettingsStruct(username, password)
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configPath, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func readSettings() (*settingsStruct, error) {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var settings settingsStruct
	err = json.Unmarshal(file, &settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}
