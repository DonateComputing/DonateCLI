package app

import (
	"github.com/mfigurski80/DonateCLI/api"
)

// ListHub returns all public hub jobs
func ListHub(auth api.AuthStruct, user bool) (api.JobListStruct, error) {
	var jobs api.JobListStruct
	if user {
		user, err := api.GetUser(auth)
		if err != nil {
			return nil, err
		}
		jobs = make(api.JobListStruct, len(user.Authored))
		i := 0
		for _, v := range user.Authored {
			jobs[i] = v
			i++
		}
	} else {
		jobsRef, err := api.GetJobs()
		if err != nil {
			return nil, err
		}
		jobs = *jobsRef
	}

	filteredJobs := filterJobs(jobs, func(j api.JobStruct) bool {
		if user && j.Author != auth.Username {
			return false
		}
		if j.CompletedImage != "" {
			return false
		}
		return true
	})

	return filteredJobs, nil
}
