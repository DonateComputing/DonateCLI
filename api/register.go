package api

import "fmt"

// AuthStruct is struct required by api to register a user and to auth existing user
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// PostRegister sends a register request to api with given data
func PostRegister(data *AuthStruct) (ResponseStruct, error) {
	// set up + do request
	res, err := doRequest("POST", domain+"/register", &data, AuthStruct{})
	if err != nil {
		return ResponseStruct{}, err
	}

	// parse body
	var r ResponseStruct
	err = parseResponseBody(res, &r)
	if err != nil {
		return r, err
	}

	// check status
	if res.StatusCode != 200 || !r.Success {
		return r, fmt.Errorf("`/register` rejected (%d), '%s'", res.StatusCode, r.Message)
	}

	return r, nil
}
