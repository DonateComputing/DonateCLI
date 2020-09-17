package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// GetJobs gets data about all current public jobs
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

// AddJob creates a new job and returns its id
func AddJob(a AuthStruct, j PostJobStruct) (string, error) {
	buf, _ := json.Marshal(j)
	req, err := http.NewRequest("POST", domain+"/job", bytes.NewBuffer(buf))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(a.Username, a.Password)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", errors.New("Rejected by server: " + string(bodyBytes))
	}

	var r CreatedResourceResponse
	err = json.Unmarshal(bodyBytes, &r)
	if err != nil {
		return "", err
	}

	return r.CreatedID, nil
}

// GetJob gets data about a specific job from public hub
func GetJob(id string) (*JobStruct, error) {
	req, err := http.NewRequest("GET", domain+"/job/"+id, nil)
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

	var job JobStruct
	err = json.Unmarshal(bodyBytes, &job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

// DeleteJob removes job with given Id and all references to it
func DeleteJob(a AuthStruct, id string) error {
	req, err := http.NewRequest("DELETE", domain+"/job/"+id, nil)
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
