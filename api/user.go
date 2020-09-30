package api

import (
	"errors"
)

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

// GetUser sends a request for authenticated user's data
func GetUser(auth AuthStruct) (*UserStruct, error) {
	// do request
	res, err := doRequest("GET", urlUser(auth.Username), nil, auth)
	if err != nil {
		return nil, err
	}

	// check response
	if res.StatusCode != 200 {
		r, err := parseUpdateResponseBody(res)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(r.Message)
	}

	// return body
	var u UserStruct
	err = parseResponseBody(res, &u)
	return &u, err
}

// UpdatePassword sends request to update authenticated user's password
func UpdatePassword(newPassword string, auth AuthStruct) error {
	data := AuthStruct{
		Username: auth.Username,
		Password: newPassword,
	}
	// do request
	res, err := doRequest("PUT", urlUser(auth.Username), data, auth)
	if err != nil {
		return err
	}

	// read/check response
	r, err := parseUpdateResponseBody(res)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New(r.Message)
	}

	return nil
}

// DeleteUser sends request to delete authenticated user
func DeleteUser(auth AuthStruct) error {
	// do request
	res, err := doRequest("DELETE", urlUser(auth.Username), nil, auth)
	if err != nil {
		return err
	}

	// read/check response
	r, err := parseUpdateResponseBody(res)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New(r.Message)
	}

	return nil
}
