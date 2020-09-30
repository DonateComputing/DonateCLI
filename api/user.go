package api

import "fmt"

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
func GetUser(username string, auth AuthStruct) (UserStruct, error) {
	// do request
	res, err := doRequest("GET", fmt.Sprintf("%s/%s", domain, username), nil, auth)
	if err != nil {
		return UserStruct{}, err
	}

	// parse body
	var u UserStruct
	err = parseResponseBody(res, &u)
	if err != nil {
		var r ResponseStruct
		err = parseResponseBody(res, &r)
		if err != nil {
			return u, err
		}
		return u, makeErrorFromResponse(res, r)
	}

	return u, nil
}
