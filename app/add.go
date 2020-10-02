package app

import "github.com/mfigurski80/DonateCLI/api"

// Add creates a new job and uploads it to hub
func Add(title string, description string, image string, auth api.AuthStruct) error {
	job := api.JobPostStruct{
		Title:       title,
		Description: description,
		Image:       image,
	}
	return api.PostJob(job, auth)
}
