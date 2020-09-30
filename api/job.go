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

// JobRefStruct is struct used by api as absolute reference to specific job
type JobRefStruct struct {
	User  string `json:"user"`
	Title string `json:"title"`
}

// PostJobStruct is struct sent to api when creating a new job
type PostJobStruct struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
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

// PostJob send request to create a job from given data
func PostJob(data PostJobStruct, auth AuthStruct) error {
	// do request
	res, err := doRequest("POST", domain+"/job", data, auth)
	if err != nil {
		return err
	}

	// read response
	r, err := parseUpdateResponseBody(res)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 || !r.Success {
		return errors.New(r.Message)
	}

	return nil
}

// GetJob sends query for specific job of given ref
func GetJob(ref JobRefStruct) (*JobStruct, error) {
	// do request
	res, err := doRequest("GET", urlUserJob(ref.User, ref.Title), nil, AuthStruct{})
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

	// read response
	var job JobStruct
	err = parseResponseBody(res, &job)
	return &job, err
}
