package app

import "github.com/mfigurski80/DonateCLI/api"

// Remove given job from hub
func Remove(auth api.AuthStruct, title string) error {
	return api.DeleteJob(api.JobRefStruct{User: auth.Username, Title: title}, auth)
}
