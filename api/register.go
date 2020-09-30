package api

import "errors"

// AuthStruct is struct required by api to register a user and to auth existing user
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// PostRegister sends a register request to api with given data
func PostRegister(data AuthStruct) error {
	// set up + do request
	res, err := doRequest("POST", domain+"/register", &data, AuthStruct{})
	if err != nil {
		return err
	}

	// parse body
	r, err := parseUpdateResponseBody(res)
	if err != nil {
		return err
	}

	// respond, and check status
	if res.StatusCode != 200 {
		return errors.New(r.Message)
	}
	return nil
}
