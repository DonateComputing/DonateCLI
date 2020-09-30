package api

// UserStruct is struct returned by api when querying user data
type UserStruct struct {
	Username string         `json:"username"`
	Password uint32         `json:"password"`
	Authored JobMapStruct   `json:"authored"`
	Running  []JobRefStruct `json:"running"`
}

// JobRefStruct is struct returned by api as absolute reference to specific job
type JobRefStruct struct {
	User  string `json:"user"`
	Title string `json:"title"`
}

// GetUser sends a request for the user with given username
func GetUser(username string, auth AuthStruct) (*UserStruct, error) {

	return &UserStruct{}, nil
}
