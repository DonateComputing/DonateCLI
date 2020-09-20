package app

import (
	"github.com/mfigurski80/DonateCLI/api"
)

// ListHub returns all public hub jobs
func ListHub(auth api.AuthStruct, all bool, user bool) ([]api.JobStruct, error) {
	jobs, err := api.GetJobs()
	if err != nil {
		return nil, err
	}

	filteredJobs := filterJobs(*jobs, func(j api.JobStruct) bool {
		if user && j.Author != auth.Username {
			return false
		}
		if j.CompletedImage != "" && !all {
			return false
		}
		return true
	})

	return filteredJobs, nil
}
