package api

import "errors"

// JobStruct is struct returned by api when querying jobs
type JobStruct struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	OriginalImage  string `json:"originalImage"`
	CompletedImage string `json:"completedImage"`
	Timestamp      int64  `json:"timestamp"`
	Author         string `json:"author"`
	Runner         string `json:"runner"`
}

// JobMapStruct is struct returned by api when querying a user's jobs
type JobMapStruct map[string]JobStruct

// JobListStruct is struct returned by api when querying for free jobs
type JobListStruct []JobStruct

// GetJobs sends query for free jobs
func GetJobs() (*JobListStruct, error) {
	// do request
	res, err := doRequest("GET", domain+"/job", nil, AuthStruct{})
	if err != nil {
		return nil, err
	}

	// check response
	if res.StatusCode != 200 {
		r, err := parseUpdateResponseBody(res)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(r.Message)
	}

	// return body
	var j JobListStruct
	err = parseResponseBody(res, &j)
	return &j, err
}
