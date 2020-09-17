package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Register sends a register request
func Register(a AuthStruct) error {
	buf, _ := json.Marshal(a)
	req, err := http.NewRequest("POST", domain+"/register", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		return errors.New("Rejected by server: " + string(bodyBytes))
	}
	return nil
}

// GetUser gets all data on user
func GetUser(a AuthStruct) (*UserStruct, error) {
	req, err := http.NewRequest("GET", domain+"/user", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(a.Username, a.Password)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("Rejected by server: " + string(bodyBytes))
	}

	var user UserStruct
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUserPassword updates the user password to new given
func UpdateUserPassword(a AuthStruct, newPassword string) error {
	newUser := MakeAuthStruct(a.Username, newPassword)
	buf, _ := json.Marshal(newUser)
	req, err := http.NewRequest("PUT", domain+"/user", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(a.Username, a.Password)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		return errors.New("Rejected by server: " + string(bodyBytes))
	}

	return nil
}

// DeleteUser removes user and all related jobs and references from public hub
func DeleteUser(a AuthStruct) error {
	req, err := http.NewRequest("DELETE", domain+"/user", nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(a.Username, a.Password)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		return errors.New("Rejected by server: " + string(bodyBytes))
	}

	return nil
}
