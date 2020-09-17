package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// GetJobs returns all current public jobs
func GetJobs() (*[]JobStruct, error) {
	req, err := http.NewRequest("GET", domain+"/job", nil)
	if err != nil {
		return nil, err
	}

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

	var jobsMap map[string]JobStruct
	err = json.Unmarshal(bodyBytes, &jobsMap)
	if err != nil {
		return nil, err
	}

	jobs := make([]JobStruct, 0)
	for _, j := range jobsMap {
		jobs = append(jobs, j)
	}

	return &jobs, nil
}
