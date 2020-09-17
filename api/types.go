package api

// UserStruct is the struct returned by api to represent single user
type UserStruct struct {
	Username string   `json:"username"`
	Password uint32   `json:"password"`
	Authored []string `json:"authored"`
	Running  []string `json:"running"`
}

// JobStruct is the struct returned by api to represent single job
type JobStruct struct {
	ID                   string   `json:"id"`
	OriginalImage        string   `json:"originalImage"`
	CompletedImage       string   `json:"completedImage"`
	PartialImages        []string `json:"partialImaged"`
	Title                string   `json:"title"`
	Description          string   `json:"description"`
	Timestamp            int64    `json:"timestamp"`
	AllowMultipleRunners bool     `json:"allowMultipleRunners"`
	Author               string   `json:"author"`
	Runners              []string `json:"runner"`
}

// MakeAuthStruct creates authentication struct to pass to api
func MakeAuthStruct(username string, password string) *AuthStruct {
	return &AuthStruct{
		Username: username,
		Password: password,
	}
}

// AuthStruct is the struct containing auth data to pass to api
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// MakePostJobStruct creates job data struct to pass to api
func MakePostJobStruct(title string, desc string, og string, allow bool) *PostJobStruct {
	return &PostJobStruct{
		Title:                title,
		Description:          desc,
		OriginalImage:        og,
		AllowMultipleRunners: allow,
	}
}

// PostJobStruct is the struct containing job data to pass to api
type PostJobStruct struct {
	Title                string `json:"title"`
	Description          string `json:"description"`
	OriginalImage        string `json:"originalImage"`
	AllowMultipleRunners bool   `json:"allowMultipleRunners"`
}

// CreatedResourceResponse is struct recieved when creating a resource
type CreatedResourceResponse struct {
	Message   string `json:"message"`
	CreatedID string `json:"createdId"`
}

// CheckedResourceResponse is struct recieved when taking/returning job
type CheckedResourceResponse struct {
	Message   string `json:"message"`
	CheckedID string `json:"checkedId"`
}
