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

// JobPostStruct is struct sent to api when creating a new job
type JobPostStruct struct {
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
func PostJob(data JobPostStruct, auth AuthStruct) error {
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

// DeleteJob sends request to delete specific job of given ref
func DeleteJob(ref JobRefStruct, auth AuthStruct) error {
	// do request
	res, err := doRequest("DELETE", urlUserJob(ref.User, ref.Title), nil, auth)
	if err != nil {
		return err
	}

	// read/check response
	r, err := parseUpdateResponseBody(res)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 || !r.Success {
		return errors.New(r.Message)
	}

	return nil
}

// TakeJob sends request to mark specific given job as being run
func TakeJob(ref JobRefStruct, auth AuthStruct) error {
	// do request
	res, err := doRequest("PUT", urlUserJob(ref.User, ref.Title)+"/take", nil, auth)
	if err != nil {
		return err
	}

	// read/check response
	r, err := parseUpdateResponseBody(res)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 || !r.Success {
		return errors.New(r.Message)
	}

	return nil
}

// JobReturnStruct is struct given to api when returning a job
type JobReturnStruct struct {
	Image string `json:"image"`
}

// ReturnJob sends request to mark specific given job as no longer being run
func ReturnJob(ref JobRefStruct, data JobReturnStruct, auth AuthStruct) error {
	// do request
	res, err := doRequest("PUT", urlUserJob(ref.User, ref.Title)+"/return", data, auth)
	if err != nil {
		return err
	}

	// read/check response
	r, err := parseUpdateResponseBody(res)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 || !r.Success {
		return errors.New(r.Message)
	}

	return nil
}
