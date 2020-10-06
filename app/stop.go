package app

import "github.com/mfigurski80/DonateCLI/api"

// Stop returns given job and stops its execution
func Stop(user string, title string, auth api.AuthStruct) error {
	ref := api.JobRefStruct{Title: title, User: user}

	err := api.ReturnJob(ref, api.JobReturnStruct{}, auth)
	if err != nil {
		return err
	}

	return nil
}

// StopAll returns all given jobs and stops all execution
func StopAll(auth api.AuthStruct) error {
	list, err := List(auth, false)
	if err != nil {
		return err
	}

	for _, item := range list {
		err := Stop(item.User, item.Title, auth)
		if err != nil {
			return err
		}
	}
	return nil
}
