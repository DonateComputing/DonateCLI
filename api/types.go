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

// AuthStruct is the struct containing auth data to pass to api
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// MakeAuthStruct creates authentication struct to pass to api
func MakeAuthStruct(username string, password string) *AuthStruct {
	return &AuthStruct{
		Username: username,
		Password: password,
	}
}
