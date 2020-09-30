package api

func ensureUserExists(a AuthStruct) {
	RegisterUser(a)
}

func ensureUserDestroyed(a AuthStruct) {
	DeleteUser(a)
}

func ensureJobExists(j JobPostStruct, a AuthStruct) {
	PostJob(JobPostStruct{j.Title, "", ""}, auth)
}

func ensureJobDestroyed(j JobPostStruct, a AuthStruct) {
	DeleteJob(JobRefStruct{a.Username, j.Title}, auth)
}
